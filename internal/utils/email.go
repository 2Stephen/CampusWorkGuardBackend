package utils

import (
	"CampusWorkGuardBackend/internal/initialize"
	"fmt"
	"github.com/go-gomail/gomail"
)

// SendEmailCode 发送验证码
func SendEmailCode(to string, code string) error {
	c := initialize.AppConfig.Email
	// ====== 构建邮件 ======
	m := gomail.NewMessage()
	m.SetHeader("From", "CampusWorkGuard <"+c.User+">")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "CampusWorkGuard验证码")
	m.SetBody("text/html", fmt.Sprintf("你的验证码为：<b>%s</b>，有效期 5 分钟。", code))

	// ====== 3. 发送邮件 ======
	d := gomail.NewDialer(c.Host, c.Port, c.User, c.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("发送邮件失败: %v", err)
	}

	return nil
}
