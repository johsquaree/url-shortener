package main

import (
	"log"
	"net/http"

	"github.com/johsquaree/url-shortener/pkg/handlers"
	"github.com/johsquaree/url-shortener/pkg/storage"
)

func main() {
	redisStorage := storage.NewRedisStorage("localhost:6379", "", 0)
	handlers.SetStorage(redisStorage)

	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/redirect/", handlers.RedirectURL)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Server is starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
