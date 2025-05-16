package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zemld/TextAnalyzer/file-analyzer/handlers"
)

func main() {
	router := chi.NewRouter()
	// TODO: здесь надо передавать сам файл, который мы будем анализировать.
	router.Get("/files/analyze/{id}", handlers.AnalyzeFileHandler)
	router.Get("/files/wordcloud/{id}", handlers.WordCloudHandler)

	http.ListenAndServe(":8080", nil)
}
