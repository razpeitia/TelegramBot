package main

import (
    "io"
    "net/http"
    "log"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func main() {
    log.Printf("Getting config file")
    var config = getConfig()
    log.Print(config)
    

    http.HandleFunc("/", hello)
    http.ListenAndServe(":9200", nil)
}
