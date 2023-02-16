package response

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	OK             = 200
	FAIL           = 10001
	ERROR          = 500
	NOTFOUND       = 404
	AUTHENTICATION = 10000
	ParamsError    = 10001
	MatherError    = 10002
	ExpireError    = 10003
	AUTHORIZATION  = 10004
	BAN            = 10005
)

var OK_RESULT = Result{Code: OK, Message: "成功"}

var FAIL_RESULT = Result{Code: FAIL, Message: "失败"}

func SUCCESS(data interface{}) Result {
	ok := OK_RESULT
	ok.Data = data
	return ok
}

func FAILURE(code int, message string) Result {
	return Result{Code: code, Message: message}
}
