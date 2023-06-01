package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	AMap  AMap // 这里配置对应配置文件的结构即可
	Graph Graph
	Vzoom Vzoom
	Nats  Nats
}

type AMap struct {
	Key string
}

type Graph struct {
	Neo4j Neo4j
}

type Neo4j struct {
	Activate bool
	Uri      string
	Username string
	Password string
}

type Vzoom struct {
	Sftp Sftp
}

type Sftp struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Nats struct {
	Activate bool
	Uri      string
}
