package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(GetPlayers)
	log.Fatal(http.ListenAndServe(":3333", handler))
}
