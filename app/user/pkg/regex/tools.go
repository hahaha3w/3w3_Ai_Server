package regex

import (
	"regexp"
	"strings"
)

// 定义正则表达式常量
const (
	UsernameRegex = `^[a-zA-Z0-9_-]{2,24}$`
	PasswordRegex = `^.{6,16}$`
	PhoneRegex    = `^1([38][0-9]|4[579]|5[0-35-9]|6[6]|7[0135678]|9[89])\d{8}$`
	EmailRegex    = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+$`
)

// 编译正则表达式并存储在变量中以供复用
var (
	regexUsername = regexp.MustCompile(UsernameRegex)
	regexPassword = regexp.MustCompile(PasswordRegex)
	regexPhone    = regexp.MustCompile(PhoneRegex)
	regexEmail    = regexp.MustCompile(EmailRegex)
)

// IsUsernameInvalid 检查用户名是否无效
func IsUsernameInvalid(username string) bool {
	return !isValid(regexUsername, username)
}

// IsPasswordInvalid 检查密码是否无效
func IsPasswordInvalid(password string) bool {
	return !isValid(regexPassword, password)
}

// IsPhoneInvalid 检查手机号是否无效
func IsPhoneInvalid(phone string) bool {
	return !isValid(regexPhone, phone)
}

// IsEmailInvalid 检查邮箱是否无效
func IsEmailInvalid(email string) bool {
	return !isValid(regexEmail, email)
}

// isValid 辅助函数，用于检查字符串是否匹配给定的正则表达式
func isValid(re *regexp.Regexp, str string) bool {
	if strings.TrimSpace(str) == "" {
		return false // 或者根据需求返回 true 表示空字符串无效
	}
	return re.MatchString(str)
}
