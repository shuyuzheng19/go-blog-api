package common

import "time"

// 爬虫请求头
const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.53"

// 每页获取文件的条数
const FileCount = 15

// 服务器文件路径
const SERVER_CONFIG_PATH = "/zsy/myBlog/config/config.ini"

// 服务器字体文件
const SERVER_FONT_CONFIG_PATH = "/zsy/myBlog/config/comic.ttf"

// redis key键
const (
	//存储邮箱验证码的KEY
	ValidateCode = "VALIDATE-EMAIL-CODE"

	//存储登录验证码的KEY
	LoginValidateCode = "VALIDATE-LOGIN-CODE"

	//存储用户TOKEN
	UserToken = "USER-TOKEN"

	//存储用户
	User = "USER"

	//热门博客
	HotBlog = "HOT-BLOG"

	//热门博客多久更新一次
	HotBlogExpire = time.Minute * 30

	//博客信息
	BlogMap = "BLOG-MAP"

	//分类信息
	CategoryList = "CATEGORY-LIST"

	//分类多久更新一次
	CategoryListExpire = time.Hour * 3

	//随机博客
	RandomBlog = "RANDOM-BLOG-IDS"

	//随机标签
	RandomTag = "RANDOM-TAG"

	//随机获取标签的数量
	RandomTagCount = 25

	//音乐播放列表
	MusicPlayList = "MUSIC-PLAY-LIST"

	//存放博客浏览量
	EyeCountMap = "BLOG-EYE-COUNT-MAP"

	//存放博客点赞量
	LikeCountMap = "BLOG-LIKE-COUNT-MAP"

	//用户的点赞存储
	UserLikeSet = "USER-LIKED"

	//存储用户今天所浏览的博客
	UserTodayEyeBlog = "USER-TODAY-EYE"

	//用户保存的博客
	UserSaveBlogMap = "USER-SAVE-BLOG-MAP"

	//推荐博客
	RECOMMEND = "RECOMMEND-BLOG"
)

// ES
const (
	BlogIndex = "blogs"

	FileIndex = "files"
)

// 大小换算
const (
	KB = 1024
	MB = 1024 * 1024
	GB = 1024 * 1024 * 1024
	TB = 1024 * 1024 * 1024 * 1024
)

// 数据
const (
	//博客一页多少条
	PageSize = 10
	//博客归档一页多少条
	ArchivePageSize = 30
	//每页多少条专题数据
	TopicSize = 30
	//相关文章多少条数据
	SimilarBlogCount = 15
)

// Roles
const (
	RoleAdmin      = "ADMIN"
	RoleUser       = "USER"
	RoleSuperAdmin = "SUPER_ADMIN"
)

// 上传相关参数
var (
	/*=================================上传文件参数=================================*/

	//上传文件的路径
	UploadPath = "/zsy/myBlog/app/static"

	//上传文件最大大小
	MaxUploadSize = GB * 0.5

	//所允许上传图片的格式
	ImageTypes = []string{"png", "jpg", "gif", "ico", "jpeg"}
)

// token相关参数
var (
	//token 失效时间
	TokenExpire = time.Hour * 24 * 7
	//token 加密key
	TokenEncrypt = "xiaoyuzuishuaila!!!!!!!!!"
	//Token请求头名
	TokenHeader = "Authorization"
	//Token前缀
	TokenPrefix = "Bearer "
)
