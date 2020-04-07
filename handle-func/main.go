package main

import "fmt"

type HandlerFunc func()

func (f HandlerFunc) serveHttp() {
	f()
}

type Handler interface {
	serveHttp()
}

func test(handler Handler) {
	handler.serveHttp()
}

type Cow struct {
	name string
}

func (c Cow) serveHttp() {
	fmt.Print("cow serveHttp")
}

func main() {
	test(Cow{"Max"})
	test(HandlerFunc(func() {
		fmt.Print("ano func")
	}))
}
