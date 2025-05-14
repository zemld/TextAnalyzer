package main

import (
	"net/http"

	"github.com/zemld/TextAnalyzer/core/handlers"
)

func main() {
	http.HandleFunc("/upload", handlers.UploadFileHandler)
	http.HandleFunc("/download", handlers.DownloadFileHandler)
	http.HandleFunc("/analyze", handlers.AnalyzeFileHandler)
	http.HandleFunc("/wordcloud", handlers.WordCloudHandler)

	http.ListenAndServe(":8080", nil)
}
