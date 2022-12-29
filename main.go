package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sodiumlabs/wyre-service/service"
)

func main() {
	port := os.Getenv("port")

	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("."))
	})

	svc := service.NewWyreServiceWithENV()
	webrpcHandler := service.NewWyreServiceServer(
		svc,
	)
	r.Handle("/*", webrpcHandler)
	fmt.Printf("listen :%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
