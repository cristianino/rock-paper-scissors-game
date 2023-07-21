package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	port := ":8080"

	//Router instance
	router := http.NewServeMux()

	//statics files
	fs := http.FileServer(http.Dir("public"))

	router.Handle("/public/", http.StripPrefix("/public/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	log.Printf("Listen tcp %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
