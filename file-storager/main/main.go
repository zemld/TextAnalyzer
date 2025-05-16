package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/zemld/TextAnalyzer/file-storager/handlers"
)

func main() {
	router := chi.NewRouter()

	router.Post("/files/upload", handlers.UploadFileHandler)
	router.Get("/files/{id}", handlers.GetFileHandler)
	router.Post("/files/analysis/{id}", handlers.SaveAnalysisResultHandler)
	router.Get("/files/analysis/{id}", handlers.GetAnalysisResultHandler)
}
