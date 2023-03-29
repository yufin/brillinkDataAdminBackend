package file_store

import (
	"testing"
)

func TestOSSUpload(t *testing.T) {
	e := OXS{"oss-cn-shanghai.aliyuncs.com", "LTAI4GEreWnQ5xaYTFcUkG7e", "Wp3yx4ltjAxrHQLkKQTf84qDgYET5p", "zxynt"}
	var oxs = e.Setup(AliYunOSS)
	err := oxs.UpLoad("test.png", "./test.png")
	if err != nil {
		t.Error(err)
	}
	t.Log("ok")
}
