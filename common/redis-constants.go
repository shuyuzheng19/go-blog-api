package common

import "time"

/**======================用户缓存======================**/

const (
	USER_INFO_KEY        = "USER-INFO:"
	EMAIL_CODE_KEY       = "EMAIL-CODE:"
	LOGIN_IMAGE_CODE_KEY = "LOGIN-IMAGE-CODE:"
	TOKEN_KEY            = "USER-TOKEN:"
	BLOG_CONFIG          = "BLOG-CONFIG"
)

const (
	USER_INFO_EXPIRE        = time.Minute * 30
	EMAIL_CODE_EXPIRE       = time.Minute * 2
	LOGIN_IMAGE_CODE_EXPIRE = time.Minute * 1
	TOKEN_EXPIRE            = time.Hour * 24 * 30
	BLOG_CONFIG_EXPIRE      = time.Duration(-1)
)

/**======================博客缓存======================**/

const (
	HOT_BLOG_KEY             = "HOT-BLOG"
	RECOMMEND_BLOG_KEY       = "RECOMMEND_BLOG"
	RANDOM_BLOG_KEY          = "RANDOM-BLOG"
	FIRST_PAGE_BLOG_PAGE_KEY = "FIRST-PAGE-BLOG"
	SEARCH_KEYWORD_KEY       = "SEARCH-KEYWORD"
	BLOG_INFO_MAP_KEY        = "BLOG-MAP"
	BLOG_EYE_COUNT_MAP_KEY   = "BLOG-EYE-COUNT-MAP"
	BLOG_LIKE_COUNT_MAP_KEY  = "BLOG-LIKE-COUNT-MAP"
	USER_EDITOR_SAVE_MAP     = "USER-EDITOR-SAVE-MAP"
)

const (
	HOT_BLOG_EXIPRE        = time.Hour * 1
	RECOMMEND_BLOG_EXPIRE  = time.Duration(-1)
	RANDOM_BLOG_EXPIRE     = time.Hour * 24 * 7
	FIRST_BLOG_PAGE_EXIPRE = time.Hour * 10
)

/**======================标签缓存======================**/

const (
	RANDOM_TAG_KEY          = "RANDOM-TAG"
	TAG_MAP_KEY             = "TAG-MAP"
	FIRST_TAG_BLOG_PAGE_KEY = "FIRST-TAG-BLOG-PAGE:"
)

const (
	RANDOM_TAG_EXPIRE     = time.Hour * 24 * 7
	FIRST_TAG_BLOG_EXPIRE = time.Hour * 5
)

/**======================分类缓存======================**/
const (
	CATEGORY_LIST_KEY = "CATEGORY-LIST"
)

const (
	CATEGORY_LIST_EXPIRE = time.Hour * 24 * 7
)

/**======================专题缓存======================**/
const (
	FIRST_PAGE_TOPIC_KEY      = "FIRST-TOPIC-PAGE"
	TOPIC_MAP_KEY             = "TOPIC-MAP"
	FIRST_TOPIC_BLOG_PAGE_KEY = "FIRST-TOPIC-BLOG-PAGE:"
)

const (
	FIRST_PAGE_TOPIC_EXIPRE      = time.Hour * 10
	TOPIC_MAP_EXPIRE             = time.Hour * 24 * 7
	FIRST_TOPIC_BLOG_PAGE_EXPIRE = time.Hour * 5
)

/**======================评论缓存======================**/
const (
	COMMENT_USER_LIKE = "USER-LIKED-COMMENTS:"
)
