package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/zemld/TextAnalyzer/file-storager/handlers"
)

// @title File Storage Manager
// @version 1.0
// @description Service for managing stored files in DB.
// @host file-storager-service:8083
// @BasePath /
func main() {
	router := chi.NewRouter()

	router.Get("/files/exists/{hash}", handlers.CheckFileExistsHandler)
	router.Post("/files/upload", handlers.UploadFileHandler)
	router.Get("/files/{id}", handlers.GetFileHandler)
	router.Post("/files/analysis/{id}", handlers.SaveAnalysisResultHandler)
	router.Get("/files/analysis/{id}", handlers.GetAnalysisResultHandler)
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://file-storager-service:8083/docs/swagger.json")))

	http.ListenAndServe(":8083", router)
}
