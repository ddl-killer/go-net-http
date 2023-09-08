package main

import (
	"io"
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong!")
	//fmt.Fprintln(w, "pong!")
}

func main() {
	http.HandleFunc("/ping", PingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
