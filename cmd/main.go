package main

import (
	"fmt"
	"go-svelte-subscriber/server"
	"log"
)

func main() {
	fmt.Printf("Starting app")

	svr := server.NewHTTPServer(":3000")
	log.Fatal(svr.ListenAndServe())
}
