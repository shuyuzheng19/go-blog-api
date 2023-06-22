package common

import "time"

/**======================用户缓存======================**/

var (
	USER_INFO_KEY        = "USER-INFO:"
	EMAIL_CODE_KEY       = "EMAIL-CODE:"
	LOGIN_IMAGE_CODE_KEY = "LOGIN-IMAGE-CODE:"
	TOKEN_KEY            = "USER-TOKEN:"
)

var (
	USER_INFO_EXPIRE        = time.Minute * 30
	EMAIL_CODE_EXPIRE       = time.Minute * 2
	LOGIN_IMAGE_CODE_EXPIRE = time.Minute * 1
	TOKEN_EXPIRE            = time.Hour * 24 * 30
)

/**======================博客缓存======================**/

var (
	HOT_BLOG_KEY         = "HOT-BLOG"
	RECOMMEND_BLOG_KEY   = "RECOMMEND_BLOG"
	RANDOM_BLOG_KEY      = "RANDOM-BLOG"
	FIREST_PAGE_BLOG_KEY = "FIRST-PAGE-BLOG"
)

var (
	HOT_BLOG_EXIPRE         = time.Hour * 1
	RECOMMEND_BLOG_EXPIRE   = time.Duration(-1)
	RANDOM_BLOG_EXPIRE      = time.Hour * 24 * 7
	FIREST_BLOG_PAGE_EXIPRE = time.Hour * 3
)
