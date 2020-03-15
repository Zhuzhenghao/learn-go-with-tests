package main

import "fmt"

// Hello sss
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Printf(Hello("Chris"))
}
