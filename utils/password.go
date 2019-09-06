package utils

import (
	"github.com/liu578101804/go-tool/encryption"
)

func PasswordCreate(password,salt string) (string,string) {
	if salt == "" {
		salt = RandString(32)
	}
	//md5(sha1(pass)+salt)
	sha1Pass := encryption.Sha1Encrypt([]byte(password))
	md5Pass := encryption.Md5Encrypt([]byte(sha1Pass + salt))
	return md5Pass,salt
}
