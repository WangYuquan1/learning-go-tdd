package main

import "fmt"

const engilishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return engilishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world"))
}
