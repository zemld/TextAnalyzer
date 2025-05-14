package main

import (
	"net/http"

	"github.com/zemld/TextAnalyzer/file-analyzer/handlers"
)

func main() {
	http.HandleFunc("/analyze", handlers.AnalyzeFileHandler)
	http.HandleFunc("/wordcloud", handlers.WordCloudHandler)

	http.ListenAndServe(":8080", nil)
}
