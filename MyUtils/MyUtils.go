package MyUtils

import (
	"crypto/md5"

	"github.com/google/uuid"
)

func Md5(str string) string {
	b := md5.Sum([]byte(str))
	return string(b[:])
}

func Uuid() string {
	return uuid.NewString()
}
