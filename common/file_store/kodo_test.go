package file_store

import (
	"testing"
)

func TestKODOUpload(t *testing.T) {
	e := OXS{"", "rMRHfXT0ieNrHYqKUqOPJQ3npxrlIPZJfsjHPKGE", "C-045VsMiJeRCZQIgZidg7IzDc5HGi1yazFuRwSK", "go-aka"}
	var oxs = e.Setup(QiNiuKodo, map[string]interface{}{"Zone": "华东"})
	err := oxs.UpLoad("test.png", "./test.png")
	if err != nil {
		t.Error(err)
	}
	t.Log("ok")
}

func TestKODOGetTempToken(t *testing.T) {
	e := OXS{"", "rMRHfXT0ieNrHYqKUqOPJQ3npxrlIPZJfsjHPKGE", "C-045VsMiJeRCZQIgZidg7IzDc5HGi1yazFuRwSK", "go-aka"}
	var oxs = e.Setup(QiNiuKodo, map[string]interface{}{"Zone": "华东"})
	token, _ := oxs.GetTempToken()
	t.Log(token)
	t.Log("ok")
}
