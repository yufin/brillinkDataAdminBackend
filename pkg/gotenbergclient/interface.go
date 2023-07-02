package gotenbergclient

type GtbCli interface {
	getUri() (string, error)
	Printing(printing GtbPrinting) ([]byte, error)
}

type GtbPrinting interface {
	targetUrl() string
}
