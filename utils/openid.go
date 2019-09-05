package utils

import (
	"time"
	"github.com/liu578101804/go-tool/encryption"
	"fmt"
)

func CreateOpenId() string {
	//MD5 16
	return encryption.Md5Encrypt([]byte(fmt.Sprintf("open-id-%v",time.Now().Unix())))[8:24]
}