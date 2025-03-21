package email

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// SendVerificationCode 发送验证码邮件
func SendVerificationCode(toEmail, code string) error {
	// 从配置文件中读取邮件配置
	user := viper.GetString("email.user")
	host := viper.GetString("email.host")
	port := viper.GetInt("email.port")
	password := viper.GetString("email.password")
	projectName := viper.GetString("email.project_name")

	// 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s (%s)", user, projectName)) // 发件人
	m.SetHeader("To", toEmail)                                     // 收件人
	m.SetHeader("Subject", fmt.Sprintf("您本次的验证码为 %s", code))       // 邮件主题

	// 设置邮件内容（HTML 格式）
	mailText := fmt.Sprintf(`
		<html>
			<body>
				<h2>尊敬的用户：</h2>
				<p>感谢您使用 <strong>%s</strong> 平台！</p>
				<p>您的验证码为：<strong style='font-size: 24px; color: #ff0000;'>%s</strong></p>
				<p>验证码将在<strong>十分钟</strong>后失效，请尽快进行验证，并不要透露给他人。</p>
				<p>祝您生活愉快！</p>
			</body>
		</html>
	`, projectName, code)
	m.SetBody("text/html", mailText)

	// 配置邮件服务器
	d := gomail.NewDialer(host, port, user, password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
