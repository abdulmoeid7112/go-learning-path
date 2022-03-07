package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello server Go!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Route: test!")
	})
	http.ListenAndServe(":8000", nil)
}
