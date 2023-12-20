package main

import "fmt"

const englishPrefix = "Hello,"
const spanishPrefix = "Hola,"
const frenchPrefix = "Bonjour,"

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := GreetingPrefix(lang)

	return fmt.Sprintf("%s %s", prefix, name)
}

func GreetingPrefix(lang string) string {
	prefix := englishPrefix
	switch lang {
	case "spanish":
		prefix = spanishPrefix
	case "french":
		prefix = frenchPrefix
	}
	return prefix
}
func main() {
	fmt.Println(Hello("Bob", ""))
}
