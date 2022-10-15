package main

// based on https://gobyexample.com/http-servers

import (
	"fmt"
	"golang.org/x/exp/maps"
	"net/http"
	"sort"
)

func headers(w http.ResponseWriter, req *http.Request) {
	keys := maps.Keys(req.Header)
	sort.Strings(keys)
	fmt.Fprintf(w, "Welcome to go-kube-example\n\n")
	for _, key := range keys {
		headers := req.Header[key]
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", key, h)
		}
	}
}

var port = 8080

func main() {
	http.HandleFunc("/", headers)
	fmt.Printf("Listening, try http://localhost:%d/\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
