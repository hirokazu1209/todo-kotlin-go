package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/hirokazu1209/todo-kotlin-go/api/handler"
)

func main() {
	r := chi.NewRouter()

	// CORS対応（Android連携時に必須）
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})
	r.Use(c.Handler)

	// ルーティング
	r.Route("/api/tasks", func(r chi.Router) {
		r.Get("/", handler.GetTasks)
		r.Post("/", handler.CreateTask)
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
