package main

import (
    "log"
    "net/http"
    "time"
)

func main() {
    // create a server
    myServer := &http.Server{
        // set the server address
        Addr: "127.0.0.1:8080",
        // define some specific configuration
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
		Handler: &myHandler{},
    }
    // launch the server
    log.Fatal(myServer.ListenAndServe())
}

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    toSend := []byte("<html><head></head><body>Hello</hello></html>")
    _, err := w.Write(toSend)
    if err != nil {
        log.Printf("error while writing on the body %s", err)
    }
}