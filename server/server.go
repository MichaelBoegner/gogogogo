package main

import (
    "log"
    "net/http"
    "time"
	"os"
)

func main() {
    // create a server
    myServer := &http.Server{
        // set the server address
        Addr: "127.0.0.1:8084",
        // define some specific configuration
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        Handler:      &myHandler{},
    }

    // launch the server
    log.Fatal(myServer.ListenAndServe())
}

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    _, err := os.Open("./temp/test.go")
    if err != nil {
        log.Println("error file open ", err)
    } else {
        log.Println("file opened")
    }
}