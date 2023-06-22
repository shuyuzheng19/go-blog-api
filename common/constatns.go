package common

const (
	//发送验证码到用户的消息
	SEND_EMAIL_SUBJECT = "ZSY-BLOG 验证码"
	//生成验证码所需字体文件的路径
	IMAGE_FONT_PATH = "comic.ttf"
	//我的邮箱
	MY_EMAIL = "shuyuzheng19@gmail.com"
)

var (
	LOGIN_AUTH = "/*/auth/**"
	ADMIN_AUTH = "/*/admin/**"
	SUPER_AUTH = "/admin/**"
	USER_ROLE  = "USER"
	ADMIN_ROLE = "ADMIN"
	SUPER_ROLE = "SUPER_ADMIN"
)

const (
	//Token加密字符串
	TOKEN_ENCRYPT = "XIAOYUZUISHUAILE"
	//Token对应的请求头
	TOKEN_HEADER = "Authorization"
	//Token类型
	TOKEN_TYPE = "Bearer "
)

const (
	PAGE_SIZE        = 10
	HOT_BLOG_SIZE    = 10
	RANDOM_BLOG_SIZE = 10
)
