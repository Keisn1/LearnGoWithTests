package hello

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func Hello(r string) string {
	if r == "" {
		r = "World"
	}
	return fmt.Sprintf("%s%s", englishHelloPrefix, r)
}

func main() {
}
