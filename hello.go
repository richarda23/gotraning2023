package main

import "fmt"

const englishPrefix = "Hello,"
const spanishPrefix = "Hola,"

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishPrefix
	if lang == "spanish" {
		prefix = spanishPrefix
	}
	return fmt.Sprintf("%s %s", prefix, name)
}
func main() {
	fmt.Println(Hello("Bob", ""))
}
