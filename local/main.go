package main

// +build !appengine

import (
	"log"
	"net/http"

	_ "github.com/dgryski/reqecho"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
