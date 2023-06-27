package task

import (
	"context"
	"encoding/json"
	"fmt"
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
	"sync/atomic"
	"time"
)

const signalMergeDuplicate = "signMergeDuplicate"

var running int32

type SyncGraphTask struct {
}

func (t SyncGraphTask) Exec(arg interface{}) error {
	if atomic.LoadInt32(&running) == 1 {
		log.Info("SyncGraph任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&running, 1)
	defer atomic.StoreInt32(&running, 0)

	md := mergeGraphDuplicated{}
	defer func() {
		if err := md.mergeAll(); err != nil {
			log.Errorf("Defer exec mergeAll error: %v", err)
		}
		if err := md.connectByLabel(); err != nil {
			log.Errorf("Defer exec connectByLabel error: %v", err)
		}
	}()

	for {
		msgs, err := natsclient.SubToSyncGraphNew.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				if err := t.AssigningTask(); err != nil {
					return err
				}
				return nil
			}
			return errors.WithStack(err)
		}
		for _, msg := range msgs {
			msgStr := string(msg.Data)
			if msgStr == signalMergeDuplicate {
				if err := md.mergeAll(); err != nil {
					return err
				}
			} else {
				err := t.SyncGraph(msgStr)
				if err != nil {
					return errors.WithStack(err)
				}
			}
			if err := msg.AckSync(); err != nil {
				return errors.WithStack(err)
			}
		}
	}
}

func (t SyncGraphTask) AssigningTask() error {
	limit := 1000
	offset := 0
	sentCounter := 0
	var dtInfo modelsSp.EnterpriseInfo
	dbInfo := sdk.Runtime.GetDbByKey(dtInfo.TableName())

	for {
		uscIds := make([]string, 0)
		var temp []struct {
			InfoId int64
			UscId  string
		}
		err := dbInfo.Table(dtInfo.TableName()).
			Select("DISTINCT usc_id, info_id").
			Order("info_id desc").
			Limit(limit).
			Offset(offset).
			Scan(&temp).
			Error
		if err != nil {
			return errors.WithStack(err)
		}

		if len(temp) == 0 {
			break
		}
		for _, v := range temp {
			uscIds = append(uscIds, v.UscId)
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
				sentCounter++
				// sign trigger to merging duplicate nodes
				if sentCounter >= 2000 {
					_, err := natsclient.TaskJs.Publish(natsclient.TopicToSyncGraphNew, []byte(signalMergeDuplicate))
					if err != nil {
						return errors.WithStack(err)
					}
					sentCounter = 0
				}
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
						create (n)-[:ATTACH_TO {id: $relId}]->(c) ;`
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
						create (n)-[:ATTACH_TO{id: $relId}]->(c);`
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
						create (n)-[r:ATTACH_TO]->(c) set r += $propRel;`
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
					create (n)-[r:ATTACH_TO]->(c) set r += $propRank;`
		_, err = models.CypherWrite(context.Background(), cypherCreate, m)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

type mergeGraphDuplicated struct {
}

func (t mergeGraphDuplicated) mergeAll() error {
	log.Info("Start workflow: merge nodes and relationships")
	if err := t.mergeFlow("Product", "title"); err != nil {
		return err
	}
	if err := t.mergeFlow("Industry", "title"); err != nil {
		return err
	}
	if err := t.mergeFlow("Certification", "title"); err != nil {
		return err
	}
	if err := t.mergeFlow("RankingList", "_listId"); err != nil {
		return err
	}
	if err := t.mergeFlow("Certification", "title"); err != nil {
		return err
	}
	if err := t.mergeFlow("Company", "uscId"); err != nil {
		return err
	}
	return nil
}

func (t mergeGraphDuplicated) mergeFlow(label string, identKey string) error {
	valueDupes, err := t.getNodeIdentDupes(label, identKey)
	if err != nil {
		return err
	}
	for _, v := range valueDupes {
		v, ok := v.(string)
		if !ok {
			return errors.Errorf("Mergeflow valueDupes type assertion error: %v", v)
		}
		if err := t.mergeNodes(label, identKey, v); err != nil {
			return err
		}
		log.Infof("Complete merge duplicate with (:%s {%s:%s}", label, identKey, v)
	}
	return nil
}

func (t mergeGraphDuplicated) mergeNodes(label string, identKey string, value string) error {
	cypherMerge :=
		fmt.Sprintf(`MATCH (n:%s {%s: $value})  
						WITH n order by elementId(n) desc 
						WITH collect(n) as nodes 
						call apoc.refactor.mergeNodes(nodes, {properties: 'discard', mergeRels: true}) 
						YIELD node
						return count(*) as count;`, label, identKey)
	_, err := models.CypherWrite(context.Background(), cypherMerge, map[string]any{"value": value})
	if err != nil {
		return err
	}
	return nil
}

func (t mergeGraphDuplicated) getNodeIdentDupes(label string, identKey string) ([]any, error) {
	cypherGetDupes :=
		fmt.Sprintf(`WITH $identKey AS propertyKey
						MATCH (n: %s)
						WITH propertyKey, collect(distinct n[propertyKey]) AS values
						UNWIND values AS value
						MATCH (m:%s) 
						WHERE m[propertyKey] = value
						WITH value, collect(m) AS nodes
						where size(nodes) > 1
						RETURN collect(value) as dupes`, label, label)
	record, err := models.CypherQuery(context.Background(), cypherGetDupes, map[string]any{"identKey": identKey})
	if err != nil {
		return []any{}, err
	}
	if len(record) == 0 {
		return []any{}, nil
	}
	dupes, ok := record[0].Get("dupes")
	if !ok {
		return []any{}, errors.New("Data with Key dupes not found.")
	}
	return dupes.([]any), nil
}

func (t mergeGraphDuplicated) connectByLabel() error {
	cypherMergeApp := `merge (n:Application{title:'标签',id: '3f543cff-5d66-44e9-805f-4d3f8c27ecd2'}) 
					merge (c1:Classification {title:'榜单标签', id: 'bd7a1d12-79d2-41f8-b330-aa2d8a504f36'}) 
					merge (c2:Classification {title:'行业标签', id: '45456ba5-8da7-4a6d-8fef-5daed800c794'}) 
					merge (c3:Classification {title:'认证标签', id: 'bdb108d3-1278-480b-90d7-5a98c03240f8'}) 
					merge (c4:Classification {title:'产品标签', id: '481d52e2-7620-4448-96a1-b6a47e66bdb9'})`
	_, err := models.CypherWrite(context.Background(), cypherMergeApp, nil)
	if err != nil {
		return err
	}
	cypherConnProd := `match (c:Classification{title:'产品标签'})
					match (t:Product) where not (c)--(t) 
					create (c)-[:CLASSIFY_OF{id:randomUUID()}]->(t)`
	_, err = models.CypherWrite(context.Background(), cypherConnProd, nil)
	if err != nil {
		return err
	}
	cypherConnInd := `match (c:Classification{title:'行业标签'})
					match (t:Industry) where not (c)--(t) 
					create (c)-[:CLASSIFY_OF{id:randomUUID()}]->(t)`
	_, err = models.CypherWrite(context.Background(), cypherConnInd, nil)
	if err != nil {
		return err
	}

	cypherConnCert := `match (c:Classification{title:'认证标签'})
					match (t:Certification) where not (c)--(t) 
					create (c)-[:CLASSIFY_OF{id:randomUUID()}]->(t)`
	_, err = models.CypherWrite(context.Background(), cypherConnCert, nil)
	if err != nil {
		return err
	}
	cypherConnRank := `match (c:Classification{title:'榜单标签'})
					match (t:RankingList) where not (c)--(t) 
					create (c)-[:CLASSIFY_OF{id:randomUUID()}]->(t)`
	_, err = models.CypherWrite(context.Background(), cypherConnRank, nil)
	if err != nil {
		return err
	}
	return nil
}

// todo: 同步贸易关系
