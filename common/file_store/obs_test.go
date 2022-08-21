package file_store

import (
	"testing"
)

func TestOBSUpload(t *testing.T) {
	e := OXS{"obs.cn-east-2.myhuaweicloud.com", "UHCHATUCMKXTHHRDKN0S", "HqiTRUZUstmgtI0oxe1VRtZLCUr8HJ4VIIQrGJJK", "go-lingshi"}
	var oxs = e.Setup(HuaweiOBS)
	err := oxs.UpLoad("test.png", "./test.png")
	if err != nil {
		t.Error(err)
	}
	t.Log("ok")
}
