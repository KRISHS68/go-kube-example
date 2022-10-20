package main

// based on https://gobyexample.com/http-servers

import (
	"fmt"
	"golang.org/x/exp/maps"
	"net/http"
	"os"
	"sort"
)

func headers(w http.ResponseWriter, req *http.Request) {
	keys := maps.Keys(req.Header)
	sort.Strings(keys)
	fmt.Fprintf(w, "Welcome to go-kube-example\n\n")
	fmt.Fprintf(w, "Arguments: %+v\n\n", os.Args)
	for _, key := range keys {
		value := req.Header.Get(key)
		fmt.Fprintf(w, "%v: %v\n", key, value)
	}
}

var port = 8080

func main() {
	http.HandleFunc("/", headers)
	fmt.Printf("Listening, try http://localhost:%d/\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
