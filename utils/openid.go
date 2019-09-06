package utils

import (
	"time"
	"github.com/liu578101804/go-tool/encryption"
	"fmt"
)

func CreateOpenId() string {
	randStr := fmt.Sprintf("open-id-%v",time.Now().Unix())
	//MD5 16
	return encryption.Md5Encrypt([]byte(randStr))[8:24]
}

func CreatePasswordCode() string {
	randStr := fmt.Sprintf("password-code-%v-%s",time.Now().Unix(),RandString(6))
	return encryption.Md5Encrypt([]byte(randStr))
}