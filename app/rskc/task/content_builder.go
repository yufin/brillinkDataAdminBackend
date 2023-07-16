package task

type ContentBuilder interface {
	GetContent() *[]byte
	GetContentId() int64
	ModifyData(mdf func(*ContentBuilder, ...interface{}) error) *ContentBuilder
}
