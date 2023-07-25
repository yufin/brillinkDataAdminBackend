package v3

type Collator interface {
	Collating(*[]byte) error
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
