package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	var a int
	var b int32
	a = 15
	b = int32(a) + int32(a)
	b = b + 5

	log.Fatal(http.ListenAndServe(":8080", router))
}
