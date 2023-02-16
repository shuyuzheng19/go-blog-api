package modal

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type User struct {
	Id         int       `gorm:"primary_key"`
	Username   string    `gorm:"size:16;unique;notnull"`
	Password   string    `gorm:"notnull"`
	Nickname   string    `gorm:"size:20;notnull"`
	Email      string    `gorm:"notnull"`
	CreateTime time.Time `gorm:"notnull"`
	Icon       string    `gorm:"notnull"`
	Sex        int       `gorm:"size:1;default:0"`
	RoleId     int       `json:"role_id" gorm:"default:1"`
	Role       Role      `gorm:"foreignKey:RoleId;references:Id"`
	Deleted    bool      `gorm:"default:0"`
}

type UserVo struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	NickName   string `json:"nickName"`
	RoleName   string `json:"roleName"`
	Admin      bool   `json:"admin"`
	SuperAdmin bool   `json:"superAdmin"`
	Icon       string `json:"icon"`
}

type SimpleUserVo struct {
	Id       int    `json:"id"`
	NickName string `json:"nickName"`
}

func (user User) ToSimpleVo() SimpleUserVo {

	return SimpleUserVo{
		Id:       user.Id,
		NickName: user.Nickname,
	}
}

func (user User) ToVo() UserVo {

	var roleName = user.Role.Name

	return UserVo{
		Id:         user.Id,
		Username:   user.Username,
		NickName:   user.Nickname,
		RoleName:   roleName,
		Admin:      roleName == common.RoleAdmin || roleName == common.RoleSuperAdmin,
		SuperAdmin: roleName == common.RoleSuperAdmin,
		Icon:       user.Icon,
	}
}

type UserRegisteredRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	RawPassword string `json:"rawPassword"`
	Email       string `json:"email"`
	Sex         int    `json:"sex"`
	Code        string `json:"code"`
	NickName    string `json:"nickName"`
	Icon        string `json:"icon"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func (user UserRegisteredRequest) ToUserDo() User {

	if user.Icon == "" || strings.TrimSpace(user.Icon) == "" {
		user.Icon = "/favicon.ico"
	}

	return User{
		Id:         0,
		Username:   user.Username,
		Password:   user.Password,
		Nickname:   user.NickName,
		Email:      user.Email,
		CreateTime: time.Now(),
		Icon:       user.Icon,
		Sex:        user.Sex,
		RoleId:     0,
		Role:       Role{},
		Deleted:    false,
	}
}

// 验证用户
func (user UserRegisteredRequest) Validator() {

	if utils.IsEmpty(user.Username) {
		panic(response.GlobalException{Code: response.ParamsError, Message: "账号不能为空!"})
	}

	var usernameLen = len(user.Username)

	if usernameLen < 8 || usernameLen > 16 {
		panic(response.GlobalException{Code: response.ParamsError, Message: "账号字符不能小于8个并且不能大于16个!"})
	}

	if utils.IsEmpty(user.Password) {
		panic(response.GlobalException{Code: response.ParamsError, Message: "密码不能为空!"})
	}

	var passwordLen = len(user.Password)

	if passwordLen < 8 || passwordLen > 16 {
		panic(response.GlobalException{Code: response.ParamsError, Message: "密码字符不能小于8个并且不能大于16个!"})
	}

	if user.Password != user.RawPassword {
		panic(response.GlobalException{Code: response.ParamsError, Message: "两次密码不一致"})
	}

	if utils.IsEmpty(user.Email) {
		panic(response.GlobalException{Code: response.ParamsError, Message: "邮箱不能为空"})
	}

	if !utils.IsEmail(user.Email) {
		panic(response.GlobalException{Code: response.ParamsError, Message: "这不是一个正确的邮箱格式"})
	}

	if user.Sex != 0 && user.Sex != 1 {
		panic(response.GlobalException{Code: response.ParamsError, Message: "性别参数有误"})
	}

	if utils.IsEmpty(user.Code) {
		panic(response.GlobalException{Code: response.ParamsError, Message: "验证码不能为空"})
	}

	if len(user.Code) != 6 {
		panic(response.GlobalException{Code: response.ParamsError, Message: "验证码格式错误,验证码为6位数字"})
	}

	if utils.IsEmpty(user.NickName) {
		panic(response.GlobalException{Code: response.ParamsError, Message: "用户名称不能为空"})
	}

}

func GetUser(ctx *gin.Context) User {
	result, exists := ctx.Get("user")

	if !exists {
		panic(response.NewGlobalException(response.AUTHENTICATION, "认证失败"))
	}

	user := result.(User)

	return user
}
