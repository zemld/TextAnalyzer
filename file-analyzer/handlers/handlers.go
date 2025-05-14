package handlers

import (
	"net/http"
)

func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проводим его анализ.
	// Отдаем результат анализа на выход.
}

func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проводим облако слов.
	// Отдаем картинку с облаком на выход.
}
