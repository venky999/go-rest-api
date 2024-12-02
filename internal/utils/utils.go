package utils

// call panic
func PanicIfError(err interface{}) {
	if err != nil {
		panic(err)
	}
}
