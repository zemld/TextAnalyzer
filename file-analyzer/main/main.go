package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zemld/TextAnalyzer/file-analyzer/handlers"
)

// @title File Analyzer
// @version 1.0
// @description Service for analyzing files.
// @host localhost:8083
// @BasePath /
func main() {
	router := chi.NewRouter()
	router.Get("/files/analyze", handlers.AnalyzeFileHandler)
	router.Get("/files/wordcloud", handlers.WordCloudHandler)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8083/docs/swagger.json")))

	http.ListenAndServe(":8083", router)
}
