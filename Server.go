package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/replace", createHandler)
	http.HandleFunc("/retrieve", retrieveHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/stats", statsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
