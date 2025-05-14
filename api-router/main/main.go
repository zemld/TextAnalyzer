package main

import (
	"net/http"

	"github.com/zemld/TextAnalyzer/api-router/handlers"
)

func main() {
	http.HandleFunc("/upload", handlers.UploadFileHandler)
	http.HandleFunc("/download", handlers.DownloadFileHandler)
	http.HandleFunc("/analyze", handlers.AnalyzeFileHandler)
	http.HandleFunc("/wordcloud", handlers.WordCloudHandler)

	// TODO: по-хорошему здесь надо делать проверку исключений, чтобы если че программа не падала.
	http.ListenAndServe(":8080", nil)
}
