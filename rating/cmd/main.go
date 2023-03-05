package main

import (
	"log"
	"net/http"

	"github.com/zhas-off/movie-service-2/rating/internal/controller/rating"
	httphandler "github.com/zhas-off/movie-service-2/rating/internal/handler/http"
	"github.com/zhas-off/movie-service-2/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
