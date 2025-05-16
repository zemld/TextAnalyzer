package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zemld/TextAnalyzer/core/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Post("/files/upload/", handlers.UploadFileHandler)
	router.Get("/files/download/{id}", handlers.DownloadFileHandler)
	router.Get("/files/analyze/{id}", handlers.AnalyzeFileHandler)
	router.Get("/files/wordcloud/{id}", handlers.WordCloudHandler)
	router.Get("/files/compare/{first-id}/{second-id}", handlers.CompareFilesHandler)

	http.ListenAndServe(":8080", nil)
}
