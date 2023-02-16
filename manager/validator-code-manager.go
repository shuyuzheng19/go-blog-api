package manager

import (
	"github.com/afocus/captcha"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/response"
)

type ValidateCodeManager struct {
}

func NewValidateManager() ValidateCodeManager {
	return ValidateCodeManager{}
}

func (*ValidateCodeManager) GenerateValidateCode(ip string) *captcha.Image {

	c := captcha.New()

	c.SetFont(common.SERVER_FONT_CONFIG_PATH)

	c.SetDisturbance(captcha.NORMAL)

	c.SetSize(100, 40)

	image, str := c.Create(4, captcha.ALL)

	config.Redis.Set(common.LoginValidateCode+":"+ip, str, time.Minute)

	return image

}

func (*ValidateCodeManager) VerifyLoginCode(currentCode string, ip string) {

	code := config.Redis.Get(common.LoginValidateCode + ":" + ip).Val()

	if code == "" {
		panic(response.NewGlobalException(response.ExpireError, "验证码可能已过期,请刷新验证码后重试"))
	}

	if currentCode != code {
		panic(response.NewGlobalException(response.NOTFOUND, "验证码错误,请核对输入"))
	}

}
