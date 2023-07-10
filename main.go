package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/text", textHandler)
	log.Println("Listening and serving on port 8080")
	http.ListenAndServe(":8080", nil)
}
