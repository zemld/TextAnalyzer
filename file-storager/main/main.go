package main

import (
	"github.com/gorilla/mux"
	"github.com/zemld/TextAnalyzer/file-storager/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/files/upload", handlers.UploadFileHandler)
	router.HandleFunc("/files/{id}", handlers.GetFileHandler)
	router.HandleFunc("/files/analysis/{id}", handlers.SaveAnalysisResultHandler).Methods("POST")
	router.HandleFunc("/files/analysis/{id}", handlers.GetAnalysisResultHandler).Methods("GET")
}
