package email

import "math/rand"

func RandomNumbers(len int) string {
	// 定义数字字符集
	digits := "0123456789"

	// 生成 len 位随机数字字符串
	randomString := make([]byte, len)
	for i := range randomString {
		randomString[i] = digits[rand.Intn(10)]
	}

	return string(randomString)
}
