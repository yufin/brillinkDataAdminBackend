package v3

import "encoding/json"

type SubjoinData struct {
	MajorCommodityProportion *float64               `json:"majorCommodityProportion"`
	IndustryTag              *[]string              `json:"industryTag"`
	AuthorizedTag            *[]AuthorizedTagDetail `json:"authorizedTag"`
	CompanyInfo              *CompanyInfoDetail     `json:"companyInfo"`
	ProductTag               *[]string              `json:"productTag"`
	RankingTag               *[]RankingTagDetail    `json:"rankingTag"`
}

type AuthorizedTagDetail struct {
	AuthClass      string `json:"authClass"`
	TagTitle       string `json:"tagTitle"`
	Authority      string `json:"authority"`
	AuthorizedDate string `json:"authorizedDate"`
}

func (e *AuthorizedTagDetail) GenMap() map[string]any {
	if e == nil {
		return nil
	}
	var m map[string]any
	b, _ := json.Marshal(*e)
	_ = json.Unmarshal(b, &m)
	return m
}

type CompanyInfoDetail struct {
	EstablishDate  string `json:"establishedDate"`
	EnterpriseType string `json:"enterpriseType"`
	CapitalPaidIn  string `json:"capitalPaidIn"`
	HomePage       string `json:"homePage"`
}

func (e *CompanyInfoDetail) GenMap() *map[string]any {
	if e == nil {
		return nil
	}
	var m map[string]any
	b, _ := json.Marshal(*e)
	_ = json.Unmarshal(b, &m)
	return &m
}

type RankingTagDetail struct {
	DatePublished string `json:"datePublished"`
	TagTitle      string `json:"tagTitle"`
	Authority     string `json:"authority"`
	Ranking       int    `json:"ranking"`
	Total         int    `json:"total"`
}

func (e *RankingTagDetail) GenMap() *map[string]any {
	if e == nil {
		return nil
	}
	var m map[string]any
	b, _ := json.Marshal(*e)
	_ = json.Unmarshal(b, &m)
	return &m
}

type SubjectCompanyTags struct {
	IndustryTag       *[]string              `json:"industryTag"`
	AuthorizedTag     *[]AuthorizedTagDetail `json:"authorizedTag"`
	ProductProportion *[]ProductProportion   `json:"productProportion"`
	RankingTag        *[]RankingTagDetail    `json:"rankingTag"`
}

func (e *SubjectCompanyTags) GenMap() *map[string]any {
	if e == nil {
		return nil
	}
	var m map[string]any
	b, _ := json.Marshal(*e)
	_ = json.Unmarshal(b, &m)
	return &m
}

type ProductProportion struct {
	Proportion     string `json:"proportion"`
	Category       string `json:"category"`
	CategoryDetail string `json:"categoryDetail"`
}

func (e *ProductProportion) GenMap() *map[string]any {
	if e == nil {
		return nil
	}
	var m map[string]any
	b, _ := json.Marshal(*e)
	_ = json.Unmarshal(b, &m)
	return &m
}
