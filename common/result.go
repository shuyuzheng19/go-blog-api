package common

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	OK_CODE        = 200
	ERROR_CODE     = 500
	FAIL_CODE      = 10001
	VALIDATOR_CODE = 10002
)

var (
	NOT_FOUNT_ERROR          = IError(404, "NOT FOUND ERROR")
	NO_LOGIN                 = IError(401, "你还未登录,请先登录")
	AUTHENTICATE_ERROR       = IError(403, "身份验证失败,重新登录试试吧")
	AUTHORIZED_ERROR         = IError(403, "授权失败,你的权限不够,请联系管理员")
	BAD_REQUEST_ERROR        = IError(400, "请求体参数绑定错误")
	REGISTRATION_FAILED      = IError(10003, "注册失败了")
	DUPLICATE_USERNAME_ERROR = IError(10004, "该账号已存在,换一个把")
	SEND_EMAIL_ERROR         = IError(10005, "发送邮件失败")
	EMAIL_WRONG_FORMAT       = IError(10006, "这不是一个正确的邮箱格式")
	EMAIL_CODE_ERROR         = IError(10007, "邮箱验证码不存在或错误")
	LOGIN_CODE_ERROR         = IError(10008, "验证码不存在或错误")
	USER_NOT_FOUNT           = IError(10009, "不存在的账号")
	PASSWORD_ERROR           = IError(10010, "密码错误,请核对输入")
	TOKEN_GENERATE_ERROR     = IError(10011, "token生成失败")
	SAVE_RECOMMEND_ERROR     = IError(10012, "推荐文章ID要大于4个,如果超出只取前4个ID")
)

func OK() Result {
	return Result{Code: OK_CODE, Message: "成功"}
}

func Success(data interface{}) Result {
	return Result{Code: OK_CODE, Message: "成功", Data: data}
}

func Fail(message string) Result {
	return Result{Code: FAIL_CODE, Message: message}
}

func IError(code int, message string) Result {
	return Result{Code: code, Message: message}
}

func Error(message string) Result {
	return Result{Code: ERROR_CODE, Message: message}
}
