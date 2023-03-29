package dto

type Personal struct {
	Name        string     `json:"name"`
	Avatar      string     `json:"avatar"`
	Userid      string     `json:"userid"`
	Email       string     `json:"email"`
	Signature   string     `json:"signature"`
	Title       string     `json:"title"`
	Group       string     `json:"group"`
	Tags        []Tag      `json:"tags"`
	NotifyCount int        `json:"notifyCount"`
	UnreadCount int        `json:"unreadCount"`
	Country     string     `json:"country"`
	Access      string     `json:"access"`
	AccessList  []string   `json:"accessList"`
	Geographic  Geographic `json:"geographic"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
	Mobile      string     `json:"mobile"`
}

type Tag struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}
type Geographic struct {
	Province Province `json:"province"`
	City     City     `json:"city"`
}

type Province struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

type City struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}
