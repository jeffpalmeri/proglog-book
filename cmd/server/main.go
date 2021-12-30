package main

import (
	"log"

	"github.com/jeffpalmeri/proglog-book/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}