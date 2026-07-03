package main

import (
	"log"
	"net/http"
)

func main() {
	Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
