package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	port := "8080"
	fmt.Printf("starting server on port %q...\n", port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func httpReadExample() {
	resp, _ := http.Get("http://wiki.com/")
	body, _ := io.ReadAll(resp.Body)
	fmt.Print("response: ")
	fmt.Println(string(body))
}
