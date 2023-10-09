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
	AMap AMap // 这里配置对应配置文件的结构即可
	//Graph      Graph
	//Vzoom      Vzoom
	//Nats       Nats
	//Minio      Minio
	//PdfConvert PdfConvert
	//PySidecar  PySidecar
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
	Sftp           Sftp
	DecisionEngine DecisionEngine
}

type Sftp struct {
	Host     string
	Port     string
	Username string
	Password string
}

type DecisionEngine struct {
	Uri string
}

type Nats struct {
	Activate bool
	Uri      string
}

type Minio struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSsl    bool
	Activate  bool
}

type PdfConvert struct {
	Gtb    Gtb
	Report Report
}

type Report struct {
	Server        string
	Path          string
	Username      string
	Password      string
	OssBucketName string
}

type Gtb struct {
	Server string
}

type PySidecar struct {
	Uri     string
	AhpPath string
}
