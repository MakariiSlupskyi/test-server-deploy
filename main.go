package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// r.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:5173"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: true,
	// 	MaxAge:           300,
	// }))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "message": "Kolya, I ❤️ u" }`))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is just a test\n"))
	})

	http.ListenAndServe(":8080", r)
}
