package manager

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
	"vs-blog-api/common"
	"vs-blog-api/config"
	"vs-blog-api/response"
	"vs-blog-api/utils"
)

type SendEmailManager struct {
	config.EmailConfig
}

func NewSendEmailManager(config config.EmailConfig) SendEmailManager {
	return SendEmailManager{config}
}

// 发送邮箱
func (c SendEmailManager) SendEmail(toEmail string, form string, subject string, randomCode string) {

	var e = email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", form, c.Username)

	e.To = []string{toEmail}

	e.Subject = subject

	e.Text = []byte("你的验证码:(" + randomCode + ") 两分钟内有效")

	err := e.Send(c.Addr, smtp.PlainAuth("", c.Username, c.Password, c.Host))

	if err != nil {
		panic(response.GlobalException{Code: response.ERROR, Message: "后台出现异常,发送邮箱失败"})
	}

	var md5Email = utils.ToMd5String([]byte(toEmail))

	config.Redis.Set(common.ValidateCode+":"+md5Email, randomCode, time.Minute*2)

}

// 验证
func (*SendEmailManager) ValidatorEmailCode(email string, code string) {

	var md5Email = utils.ToMd5String([]byte(email))

	var redisCode = config.Redis.Get(common.ValidateCode + ":" + md5Email).Val()

	fmt.Println(redisCode)

	if redisCode == "" {
		panic(response.NewGlobalException(response.NOTFOUND, "验证码可能已过期,请重试"))
	}

	if redisCode != code {
		panic(response.NewGlobalException(response.ParamsError, "验证码错误,请重试"))
	}

}
