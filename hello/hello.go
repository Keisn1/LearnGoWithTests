package hello

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Buenas Dias, "
	frenchHelloPrefix  = "Bonjour, "
)
const (
	english = "en"
	spanish = "esp"
	french  = "fr"
)

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case english:
		return englishHelloPrefix
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func Hello(r, lang string) string {
	if r == "" {
		r = "World"
	}

	prefix := greetingPrefix(lang)
	return fmt.Sprintf("%s%s", prefix, r)
}

func main() {
}
