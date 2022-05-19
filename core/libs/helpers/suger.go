package helpers

// Must 保证有返回值，若有错误直接panic
// 用法：v := Must(fn())
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
