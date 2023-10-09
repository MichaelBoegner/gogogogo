package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func getRootHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Printf("received / GET request\n")
	fmt.Fprintf(w, "This is the homepage of a website ^_^\n")
}

func getContactHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Printf("received /Contact GET request\n")
	fmt.Fprintf(w, "Conact us by filling out the form below!\n")
}

func main() {
	http.HandleFunc("/", getRootHandler)
	http.HandleFunc("/contact", getContactHandler)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server be closed")
	} else if err != nil {
		fmt.Printf("server error returned that reads: %q", err)
		os.Exit(1)
	}
}
