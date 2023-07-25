package reportbuilder

type ReportBuilder interface {
	GetReportVersion() string
	Pipeline() error
}
