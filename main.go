package main

import (
	"fmt"
	"net/http"
	"os"
	"todo-api/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ToDo-API is running"))
	})

	r.Post("/task", handlers.AddTasks)
	r.Post("/toggle/{id}", handlers.ToggleCompleted)
	r.Get("/view", handlers.ViewTasks)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":"+port, r)
}
