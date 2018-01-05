package util

// Err 判断错误是否为空
func Err(err error) {
	if err != nil {
		panic(err)
	}
}
