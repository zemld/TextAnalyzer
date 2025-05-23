package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/zemld/TextAnalyzer/file-storager/docs"
	"github.com/zemld/TextAnalyzer/file-storager/handlers"
)

// @title File Storage Manager
// @version 1.0
// @description Service for managing stored files in DB.
// @host localhost:8082
// @BasePath /
func main() {
	router := chi.NewRouter()

	router.Get("/files/exists/{id}", handlers.CheckFileExistsHandler)
	router.Post("/files/upload", handlers.UploadFileHandler)
	router.Get("/files/{id}", handlers.GetFileHandler)
	router.Post("/files/analysis", handlers.SaveAnalysisResultHandler)
	router.Get("/files/analysis/{id}", handlers.GetAnalysisResultHandler)
	router.Post("/files/wordcloud/{id}", handlers.SaveWordCloudHandler)
	router.Get("/files/wordcloud/{id}", handlers.GetWordCloudHandler)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8082/docs/swagger.json")))

	http.ListenAndServe(":8082", router)
}
