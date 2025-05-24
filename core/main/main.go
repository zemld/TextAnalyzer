package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zemld/TextAnalyzer/core/handlers"
)

// @title Core Service
// @version 1.0
// @description Core service with bussiness logic.
// @host localhost:8081
// @BasePath /
func main() {
	router := chi.NewRouter()
	router.Post("/files/upload", handlers.UploadFileHandler)
	router.Get("/files/download/{id}", handlers.DownloadFileHandler)
	router.Get("/files/analyze/{id}", handlers.AnalyzeFileHandler)
	router.Get("/files/wordcloud/{id}", handlers.WordCloudHandler)
	router.Get("/files/compare", handlers.CompareFilesHandler)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8081/docs/swagger.json")))
	http.ListenAndServe(":8081", router)
}
