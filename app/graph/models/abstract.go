package models

import "time"

type Entity interface {
	GetIdentifier() string
	GetId() interface{}
	GetTitle() string
	GetLabels() []string
	SetLabels([]string)
	GetProps() map[string]any
	SetProps(*map[string]any)
	GenMap(*map[string]any)
}

type Relation interface {
	GetIdentifier() string
	GetId() interface{}
	SetId(interface{}) interface{}
	GetSourceId() interface{}
	SetSourceId(interface{})
	GetTargetId() interface{}
	SetTargetId(interface{})
	GetLabel() string
	SetLabel(string)
	GetProps() map[string]any
	SetProps(*map[string]any)
	GenMap(*map[string]any)
}

type NodeCommon struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type EdgeCommon struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`

	SourceId any    `json:"-"`
	TargetId any    `json:"-"`
	Label    string `json:"-"`
}
