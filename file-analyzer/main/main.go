package main

import (
	"net/http"

	"github.com/zemld/TextAnalyzer/file-analyzer/handlers"
)

func main() {
	http.HandleFunc("/files/analyze/{id}", handlers.AnalyzeFileHandler)
	http.HandleFunc("/files/wordcloud/{id}", handlers.WordCloudHandler)

	http.ListenAndServe(":8080", nil)
}
