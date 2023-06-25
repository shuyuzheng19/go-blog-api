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
	NOT_FOUNT_ERROR                 = IError(404, "NOT FOUND ERROR")
	NO_LOGIN                        = IError(401, "你还未登录,请先登录")
	AUTHENTICATE_ERROR              = IError(403, "身份验证失败,重新登录试试吧")
	AUTHORIZED_ERROR                = IError(403, "授权失败,你的权限不够,请联系管理员")
	BAD_REQUEST_ERROR               = IError(400, "请求体参数绑定错误")
	REGISTRATION_FAILED             = IError(10003, "注册失败了")
	DUPLICATE_USERNAME_ERROR        = IError(10004, "该账号已存在,换一个把")
	SEND_EMAIL_ERROR                = IError(10005, "发送邮件失败")
	EMAIL_WRONG_FORMAT              = IError(10006, "这不是一个正确的邮箱格式")
	EMAIL_CODE_ERROR                = IError(10007, "邮箱验证码不存在或错误")
	LOGIN_CODE_ERROR                = IError(10008, "验证码不存在或错误")
	USER_NOT_FOUNT                  = IError(10009, "不存在的账号")
	PASSWORD_ERROR                  = IError(10010, "密码错误,请核对输入")
	TOKEN_GENERATE_ERROR            = IError(10011, "token生成失败")
	SAVE_RECOMMEND_ERROR            = IError(10012, "推荐文章ID要大于4个,如果超出只取前4个ID")
	NO_KEYWORD                      = IError(10013, "请输入关键字")
	RANGE_TIME_ERROR                = IError(10014, "开始时间不能大于结束时间")
	RANGE_TIME_EMPTY                = IError(10015, "开始或者结束时间不能为空")
	TAG_ID_ERROR                    = IError(10016, "非法的标签ID")
	TOPIC_ID_ERROR                  = IError(10017, "非法的专题ID")
	USER_ID_ERROR                   = IError(10018, "非法的用户ID")
	TAG_NOT_FOUND                   = IError(10019, "该标签不存在")
	TOPIC_NOT_FOUND                 = IError(10020, "该专题不存在")
	BLOG_NOT_FOUND                  = IError(10021, "找不到该博客")
	BLOG_ID_ERROR                   = IError(10022, "非法的博客ID")
	LIKE_BLOG_FAIL                  = IError(10023, "点赞失败")
	REPEAT_BLOG_LIKE                = IError(10024, "你已经点过赞了,无需重复点赞")
	COMMENT_ERROR                   = IError(10025, "评论失败")
	COMMENT_ID_ERROR                = IError(10026, "非法的评论ID")
	NOT_IAMGE_FILE                  = IError(10027, "这不是一张图片")
	MAX_IMAGE_SIZE_ERROR            = IError(10028, "图片文件大小超出")
	MAX_FILE_SIZE_ERROR             = IError(10029, "文件大小超出")
	OPEN_FILE_ERROR                 = IError(10030, "文件打开失败")
	UPLOAD_FILE_ERROR               = IError(10031, "文件上传失败")
	GET_UPLOAD_FILE_ERROR           = IError(10032, "获取文件上传失败")
	CHAT_TOKEN_ERROR                = IError(10033, "请携带 openai token")
	CHAT_AUTHENTICATE_ERROR         = IError(10034, "身份验证失败,请尝试用新的Token")
	RELEASE_CATEGORY_OR_TOPIC_EMPTY = IError(10035, "专题或者分类有空")
	RELEASE_TAG_EMPTY               = IError(10036, "至少要选择一个标签")
	RELEASE_BLOG_ERROR              = IError(10037, "发布博客失败")
	ADD_CATEGORY_ERROR              = IError(10038, "添加分类失败")
	ADD_TAG_ERROR                   = IError(10039, "添加标签失败")
	TAG_NAME_EMPTY_ERROR            = IError(10041, "标签名称不能为空")
	CATEGORY_NAME_EMPTY_ERROR       = IError(10042, "分类名称不能为空")
	ADD_TOPIC_ERROR                 = IError(10043, "添加专题失败")
	UPDATE_BLOG_AUTH_ERROR          = IError(10045, "只能修改自己的博客")
	UPDATE_BLOG_ERROR               = IError(10046, "修改博客失败")
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
