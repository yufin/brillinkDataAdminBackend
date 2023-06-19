package task

import (
	"context"
	"encoding/json"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	models "go-admin/app/graph/models"
	"go-admin/app/graph/service/dto"
	modelsSp "go-admin/app/spider/models"
	"go-admin/pkg/natsclient"
	"go-admin/utils"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

const signalMergeDuplicate = "signMergeDuplicate"

type SyncGraphTask struct {
}

func (t SyncGraphTask) Exec(arg interface{}) error {
	//if err := t.AssigningTask(); err != nil {
	//	return err
	//}
	for {
		msgs, err := natsclient.SubToSyncGraphNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			}
			return errors.WithStack(err)
		}
		for _, msg := range msgs {
			msgStr := string(msg.Data)
			if msgStr == signalMergeDuplicate {
				if err := MergeGraphDuplicated(); err != nil {
					return err
				}
			} else {
				err := t.SyncGraph(msgStr)
				if err != nil {
					return errors.WithStack(err)
				}
				if err := msg.AckSync(); err != nil {
					return errors.WithStack(err)
				}
			}
		}
	}
}

//type distinctUscId struct {
//	UscId string
//}

func (t SyncGraphTask) AssigningTask() error {
	limit := 1000
	offset := 0
	var dtInfo modelsSp.EnterpriseInfo
	dbInfo := sdk.Runtime.GetDbByKey(dtInfo.TableName())

	for {
		var res []struct {
			UscId     string
			CreatedAt time.Time
		}
		//uscIds := make([]string, 0)
		err := dbInfo.Table(dtInfo.TableName()).
			Select("DISTINCT usc_id, created_at").
			Order("created_at").
			Limit(limit).
			Offset(offset).
			Scan(&res).
			Error
		if err != nil {
			return errors.WithStack(err)
		}

		uscIds := make([]string, 0)
		for _, v := range res {
			uscIds = append(uscIds, v.UscId)
		}

		if len(uscIds) == 0 {
			break
		}

		cypher := `
			unwind $uscIds as value 
			OPTIONAL MATCH (n:Company) WHERE n.uscId = value
			WITH value, n 
			WHERE n IS NULL 
			RETURN collect(toString(value)) as unSynced;
		`
		unSynced, err := models.CypherQuery(
			context.Background(), cypher, map[string]any{"uscIds": uscIds})
		if err != nil {
			return errors.WithStack(err)
		}
		offset += limit

		unSyncedIds, ok := unSynced[0].Get("unSynced")

		if ok {
			for _, unSyncedId := range unSyncedIds.([]any) {
				_, err := natsclient.TaskJs.Publish(natsclient.TopicToSyncGraphNew, []byte(unSyncedId.(string)))
				if err != nil {
					return errors.WithStack(err)
				}
			}
			// sign trigger to merging duplicate nodes
			_, err := natsclient.TaskJs.Publish(natsclient.TopicToSyncGraphNew, []byte(signalMergeDuplicate))
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

func (t SyncGraphTask) SyncGraph(uscId string) error {
	var dtInfo modelsSp.EnterpriseInfo
	db := sdk.Runtime.GetDbByKey(dtInfo.TableName())

	err := db.Model(&modelsSp.EnterpriseInfo{}).
		Where("usc_id = ?", uscId).
		Order("created_at desc").
		First(&dtInfo).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	var (
		dtIndustry                        modelsSp.EnterpriseIndustry
		dtProduct                         modelsSp.EnterpriseProduct
		wg                                sync.WaitGroup
		errInd, errProd, errCert, errRank error
	)
	dtCert := make([]modelsSp.EnterpriseCertification, 0)
	dtRankList := make([]modelsSp.EnterpriseRankingList, 0)
	wg.Add(4)
	go func() {
		defer wg.Done()
		errIndTemp := db.Model(&dtIndustry).Where("usc_id = ?", uscId).Order("created_at desc").First(&dtIndustry).Error
		if errIndTemp != nil {
			if errors.Is(errIndTemp, gorm.ErrRecordNotFound) {
				errInd = nil
			} else {
				errInd = errIndTemp
			}
		}
	}()
	go func() {
		defer wg.Done()
		errProdTemp := db.Model(&dtProduct).Where("usc_id = ?", uscId).Order("created_at desc").First(&dtProduct).Error
		if errProdTemp != nil {
			if errors.Is(errProdTemp, gorm.ErrRecordNotFound) {
				errProd = nil
			} else {
				errProd = errProdTemp
			}
		}
	}()
	go func() {
		defer wg.Done()
		errCert = db.Model(&modelsSp.EnterpriseCertification{}).Raw(
			`SELECT *
			FROM (
				SELECT *, ROW_NUMBER() OVER(PARTITION BY certification_source ORDER BY created_at DESC) as rn
				FROM enterprise_certification
				WHERE usc_id = ?
			) sub
			WHERE rn = 1;`, uscId).
			Scan(&dtCert).
			Error
	}()
	go func() {
		defer wg.Done()
		errRank = db.Model(&modelsSp.EnterpriseRankingList{}).Raw(
			`select rl.*, t.usc_id, t.ranking_position, t.rank_id from
				(SELECT *
					FROM (
						SELECT *, ROW_NUMBER() OVER(PARTITION BY list_id ORDER BY created_at DESC) as rn
						FROM enterprise_ranking
						WHERE usc_id = ?
						) sub
				WHERE rn = 1) t
				left join ranking_list rl on t.list_id = rl.id`, uscId).
			Scan(&dtRankList).
			Error
	}()
	wg.Wait()
	for _, err := range []error{errInd, errProd, errCert, errRank} {
		if err != nil {
			return errors.WithStack(err)
		}
	}

	gsd := GraphSyncData{
		UscId:      uscId,
		DtCert:     dtCert,
		DtRankList: dtRankList,
		DtInd:      &dtIndustry,
		DtProd:     &dtProduct,
		DtInfo:     &dtInfo,
	}
	defer MergeGraphDuplicated()
	err = gsd.InsertGraph()
	if err != nil {
		return errors.WithStack(err)
	}
	log.Infof("SyncGraph complete for uscId = %s", uscId)
	return nil
}

type GraphSyncData struct {
	UscId      string
	DtCert     []modelsSp.EnterpriseCertification
	DtRankList []modelsSp.EnterpriseRankingList
	DtInd      *modelsSp.EnterpriseIndustry
	DtProd     *modelsSp.EnterpriseProduct
	DtInfo     *modelsSp.EnterpriseInfo
	CompanyId  string
}

func (t *GraphSyncData) InsertGraph() error {
	if err := t.syncInfo(); err != nil {
		return errors.WithStack(err)
	}
	var (
		wg                                sync.WaitGroup
		errCert, errRank, errProd, errInd error
	)
	wg.Add(4)
	go func() {
		defer wg.Done()
		errProd = t.syncProduct()
	}()
	go func() {
		defer wg.Done()
		errInd = t.syncIndustry()
	}()
	go func() {
		defer wg.Done()
		errCert = t.syncCert()
	}()
	go func() {
		defer wg.Done()
		errRank = t.syncRankList()
	}()
	wg.Wait()
	for _, err := range []error{errProd, errInd, errCert, errRank} {
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (t *GraphSyncData) syncInfo() error {
	var syncReq dto.EnterpriseInfoSyncReq
	syncReq.Assignment(t.DtInfo)

	b, err := json.Marshal(syncReq)
	if err != nil {
		return errors.WithStack(err)
	}
	propInfo := make(map[string]any)
	if err := json.Unmarshal(b, &propInfo); err != nil {
		return errors.WithStack(err)
	}
	m := map[string]any{
		"propInfo": propInfo,
	}
	cypher := `CREATE (n:Company) Set n += $propInfo;`
	_, err = models.CypherWrite(context.Background(), cypher, m)
	if err != nil {
		return errors.WithStack(err)
	}
	t.CompanyId = syncReq.Id
	return nil
}

func (t *GraphSyncData) syncProduct() error {
	if t.DtProd.UscId == "" {
		return nil
	}
	if t.CompanyId == "" {
		return errors.New("GraphSyncData.companyId Not Assigned")
	}
	products := make([]string, 0)
	err := json.Unmarshal([]byte(t.DtProd.ProductData), &products)
	if err != nil {
		log.Errorf("SyncGraphTask.productData unmarshall error on prodId = %v: %v", t.DtProd.ProdId, err)
		return nil
	}
	for _, prod := range products {
		// replace "#" with "" in prod
		prod := strings.Replace(prod, "#", "", -1)
		props := map[string]any{
			"title": prod,
			"id":    utils.NewRandomUUID(),
		}
		m := map[string]any{
			"properties": props,
			"companyId":  t.CompanyId,
			"relId":      utils.NewRandomUUID(),
		}
		cypherCreate := `CREATE (n:Tag:Product) SET n += $properties  
						WITH n 
						MATCH (c:Company{id: $companyId}) 
						create (c)-[:HAS_TAG{id: $relId}]->(n) ;`
		_, err := models.CypherWrite(context.Background(), cypherCreate, m)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (t *GraphSyncData) syncIndustry() error {
	if t.DtInd.UscId == "" {
		return nil
	}
	if t.CompanyId == "" {
		return errors.New("GraphSyncData.companyId Not Assigned")
	}
	industries := make([]string, 0)
	err := json.Unmarshal([]byte(t.DtInd.IndustryData), &industries)
	if err != nil {
		log.Errorf("SyncGraphTask.IndustryData unmarshall error on prodId = %v: %v", t.DtInd.IndId, err)
		return nil
	}
	for _, ind := range industries {
		// replace "#" with "" in prod
		ind := strings.Replace(ind, "#", "", -1)
		props := map[string]any{
			"title": ind,
			"id":    utils.NewRandomUUID(),
		}
		m := map[string]any{
			"properties": props,
			"companyId":  t.CompanyId,
			"relId":      utils.NewRandomUUID(),
		}
		cypherCreate := `CREATE (n:Tag:Industry) SET n += $properties  
						WITH n 
						MATCH (c:Company {id: $companyId}) 
						create (c)-[:HAS_TAG{id: $relId}]->(n);`
		_, err = models.CypherWrite(context.Background(), cypherCreate, m)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (t *GraphSyncData) syncCert() error {
	if t.CompanyId == "" {
		return errors.New("GraphSyncData.companyId Not Assigned")
	}

	for _, modelCert := range t.DtCert {
		var syncReq dto.EnterpriseCertificationSyncReq
		syncReq.Assignment(&modelCert)
		b, err := json.Marshal(syncReq)
		if err != nil {
			return errors.WithStack(err)
		}
		propRel := make(map[string]any)
		if err := json.Unmarshal(b, &propRel); err != nil {
			return err
		}
		propCert := map[string]any{
			"title": syncReq.CertificationTitle,
			"id":    utils.NewRandomUUID(),
		}
		m := map[string]any{
			"propRel":   propRel,
			"propCert":  propCert,
			"companyId": t.CompanyId,
			"relId":     utils.NewRandomUUID(),
		}

		cypherCreate := `CREATE (n:Tag:Certification) SET n += $propCert  
						WITH n 
						MATCH (c:Company{id: $companyId}) 
						create (c)-[r:HAS_TAG]->(n) set r += $propRel;`
		_, err = models.CypherWrite(context.Background(), cypherCreate, m)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (t *GraphSyncData) syncRankList() error {
	if t.CompanyId == "" {
		return errors.New("GraphSyncData.companyId Not Assigned")
	}

	for _, modelList := range t.DtRankList {
		var nSyncReq dto.RankingListSyncReq
		var rSyncReq dto.EnterpriseRankingSyncReq
		nSyncReq.Assignment(&modelList)
		rSyncReq.Assignment(&modelList)

		bn, err := json.Marshal(nSyncReq)
		if err != nil {
			return errors.WithStack(err)
		}
		propList := make(map[string]any)
		if err := json.Unmarshal(bn, &propList); err != nil {
			return errors.WithStack(err)
		}

		br, err := json.Marshal(rSyncReq)
		if err != nil {
			return errors.WithStack(err)
		}
		propRank := make(map[string]any)
		if err := json.Unmarshal(br, &propRank); err != nil {
			return err
		}

		m := map[string]any{
			"propList":  propList,
			"propRank":  propRank,
			"companyId": t.CompanyId,
			"relId":     utils.NewRandomUUID(),
		}

		cypherCreate := `CREATE (n:Tag:RankingList) SET n += $propList 
					WITH n 
					MATCH (c:Company{id: $companyId})
					create (c)-[r:HAS_TAG]->(n) set r += $propRank;`
		_, err = models.CypherWrite(context.Background(), cypherCreate, m)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func MergeGraphDuplicated() error {
	return nil
}
