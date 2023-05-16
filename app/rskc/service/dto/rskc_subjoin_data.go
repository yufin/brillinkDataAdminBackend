package dto

type SubjoinData struct {
	MajorCommodityProportion *float32               `json:"majorCommodityProportion"`
	IndustryTag              *[]string              `json:"industryTag"`
	AuthorizedTag            *[]AuthorizedTagDetail `json:"authorizedTag"`
	CompanyInfo              *CompanyInfoDetail     `json:"companyInfo"`
	ProductTag               *[]string              `json:"productTag"`
	RankingTag               *[]RankingTagDetail    `json:"rankingTag"`
}

type AuthorizedTagDetail struct {
	AuthClass string `json:"authClass"`
	TagTitle  string `json:"tagTitle"`
	Authority string `json:"authority"`
}

type CompanyInfoDetail struct {
	EstablishDate  string `json:"establishDate"`
	EnterpriseType string `json:"enterpriseType"`
	CapitalPaidIn  string `json:"capitalPaidIn"`
}

type RankingTagDetail struct {
	DatePublished string `json:"datePublished"`
	TagTitle      string `json:"tagTitle"`
	Authority     string `json:"authority"`
	Ranking       int    `json:"ranking"`
	Total         int    `json:"total"`
}
