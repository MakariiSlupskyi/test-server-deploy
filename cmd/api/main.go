package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type Message struct {
	UserName string `json:"userName"`
	Message  string `json:"message"`
}

func main() {
	godotenv.Load() // for dev only

	messages := []Message{}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	fmt.Print(os.Getenv("FRONTEND_URL"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "message": "Kolya, I ❤️ u" }`))
	})

	r.Get("/messages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(messages); err != nil {
			http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
			log.Printf("JSON encode error: %v", err)
		}
	})

	r.Post("/messages", func(w http.ResponseWriter, r *http.Request) {
		msg := Message{}
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
			log.Printf("JSON decode error: %v", err)
		}
		messages = append(messages, msg)
		json.NewEncoder(w).Encode(msg)
	})

	r.Delete("/messages", func(w http.ResponseWriter, r *http.Request) {
		messages = []Message{}
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
