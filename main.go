package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	port := ":8080"

	router := http.NewServeMux()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	log.Printf("Listen to %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
