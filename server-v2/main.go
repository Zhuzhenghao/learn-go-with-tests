// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	fmt.Println(server)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
