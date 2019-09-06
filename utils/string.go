package utils

import (
	"math/rand"
	"time"
)

func RandString(len int) string {
	bytes := make([]byte, len)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < len ; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
