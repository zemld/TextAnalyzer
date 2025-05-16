package main

import (
	"net/http"

	"github.com/zemld/TextAnalyzer/core/handlers"
)

func main() {
	http.HandleFunc("/files/upload/", handlers.UploadFileHandler)
	http.HandleFunc("/files/download/{id}", handlers.DownloadFileHandler)
	http.HandleFunc("/files/analyze/{id}", handlers.AnalyzeFileHandler)
	http.HandleFunc("/files/wordcloud/{id}", handlers.WordCloudHandler)
	http.HandleFunc("/files/compare/{first-id}/{second-id}", handlers.CompareFilesHandler)

	http.ListenAndServe(":8080", nil)
}
