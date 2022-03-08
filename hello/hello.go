package main

import "fmt"

const engilishHelloPrefix = "Hello, "

func Hello(name string) string {
	return engilishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world"))
}
