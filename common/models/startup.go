package models

type SystemInfo struct {
	SystemName string `json:"systemName"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

var System = new(SystemInfo)

func (*SystemInfo) Default(systemName, username, password string) {
	if systemName != "" {
		System.SystemName = systemName
	} else {
		System.SystemName = "go-admin管理系统"
	}
	if username != "" {
		System.Username = username
	} else {
		System.Username = "admin"
	}
	if password != "" {
		System.Password = password
	} else {
		System.Password = "123456"
	}
}

func (*SystemInfo) SetSystemName(systemName string) {
	if systemName != "" {
		System.SystemName = systemName
	} else {
		System.SystemName = "go-admin管理系统"
	}
}
func (*SystemInfo) SetUsername(username string) {
	if username != "" {
		System.Username = username
	} else {
		System.Username = "admin"
	}
}
func (*SystemInfo) SetPassword(password string) {
	if password != "" {
		System.Password = password
	} else {
		System.Password = "123456"
	}
}
