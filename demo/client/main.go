package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
