package v3

type Collator interface {
	Collating(content *[]byte, contentId int64) error
}

type ReportBuilderV3 struct {
	content *[]byte
}

func (s *ReportBuilderV3) GetReportVersion() string {
	return "3.0"
}

func (s *ReportBuilderV3) Pipeline() error {
	return nil
}
