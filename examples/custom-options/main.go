package main

import (
	"log"
	"net/http"

	"github.com/syntaqx/serve"
)

func main() {
	fs := serve.NewFileServer(serve.Options{
		Directory: ".",
	})
	log.Fatal(http.ListenAndServe(":8080", fs))
}
