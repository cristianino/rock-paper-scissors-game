package main

import (
	"fmt"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	port := ":8080"

	router := http.NewServeMux()
	router.HandleFunc("/", handlers.Index)

	fmt.Printf("Listen to %s\n", port)
	http.ListenAndServe(port, router)
}
