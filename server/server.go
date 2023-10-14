package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Players struct {
	name string
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	request := strings.TrimPrefix(r.URL.String(), "/players/")
	score := 0
	if request == "Mike" {
		score = 10
	}
	fmt.Fprintf(w, "player score is: %v", score)
}
