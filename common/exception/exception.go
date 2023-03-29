package exception

import "net/http"

type Exception struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	ErrCode int    `json:"errCode"`
	Err     error  `json:"err"`
}

func WithStatus(code, errCode int, err error) Exception {
	ex := Exception{Code: code, ErrCode: errCode, Err: err}
	return ex
}

func New(errCode int, err error) Exception {
	ex := Exception{Code: http.StatusOK, ErrCode: errCode, Err: err}
	return ex
}

func WithMsg(errCode int, msg string, err error) Exception {
	bl := containChineseWord(err.Error())
	if bl {
		ex := Exception{Code: http.StatusOK, ErrCode: errCode, Msg: err.Error(), Err: err}
		return ex
	} else {
		ex := Exception{Code: http.StatusOK, ErrCode: errCode, Msg: msg, Err: err}
		return ex
	}
}

// containChineseWord 判断是否存在中文
func containChineseWord(str string) bool {
	r := []rune(str)
	for _, a := range r {
		if a >= 19968 && a <= 40869 {
			return true
		}
	}
	return false
}
