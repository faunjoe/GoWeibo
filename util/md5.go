package util

import(
  	"crypto/md5"
    "encoding/hex"
    "io"
)

// 计算字符串的MD5
func GenerateStringMD5(str string) string {
  md5 := md5.New()
	io.WriteString(md5, str)
	md5Later := md5.Sum(nil)
	md5Str := hex.EncodeToString(md5Later)
  return md5Str
}
