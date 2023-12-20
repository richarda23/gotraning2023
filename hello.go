package main

import "fmt"

const englishPrefix = "Hello,"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s %s", englishPrefix, name)
}
func main() {
	fmt.Println(Hello("Bob"))
}
