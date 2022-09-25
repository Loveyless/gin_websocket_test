package utils

import "crypto/md5"

func Md5(str string) string {
	b := md5.Sum([]byte(str))
	return string(b[:])
}
