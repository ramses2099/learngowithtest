package main

import (
	"fmt"
)

const englishhelloPrefix = "Hello, "
const spanishhelloPrefix = "Hola, "
const frenchhelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchhelloPrefix
	case spanish:
		prefix = spanishhelloPrefix
	default:
		prefix = englishhelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chrirs", "Spanish"))
}
