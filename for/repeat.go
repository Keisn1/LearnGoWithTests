package iteration

import ()

const repeatCount = 5

func Repeat(char string) string {
	var res string
	for i := 0; i < repeatCount; i++ {
		res += char
	}
	return string(res)
}
