package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "anselbrandt.dev\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func api(w http.ResponseWriter, req *http.Request) {

	resp, err := http.Get("https://anselbrandt.com/api")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func main() {
	fmt.Println("server is running...")

	http.HandleFunc("/", hello)
	http.HandleFunc("/api", api)

	http.ListenAndServe(":8080", nil)
}
