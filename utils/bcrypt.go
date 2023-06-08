package utils

//密码加解密
import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// GetHashPwd 加密密码
func GetHashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePwd 验证密码
func ComparePwd(hashedPwd string, pwd string) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
