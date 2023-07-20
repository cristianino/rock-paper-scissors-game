package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "RPC web")
	})

	fmt.Printf("Listen to %s", port)
	http.ListenAndServe(port, nil)
}
