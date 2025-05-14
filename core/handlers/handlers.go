package handlers

import (
	"net/http"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход файл. Считаем хэш от него и перекидываем запрос на file-storager.
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем есть ли файл в базе данных и перекидываем запрос на file-storager.
}

func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем есть ли файл в базе данных.
	// Смотрим, есть ли уже результат анализа. Если есть, то отправляем его на выход. Если нет, то перекидываем запрос на file-analyzer.
}

func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем есть ли файл в базе данных.
	// Смотрим, есть ли уже результат облако. Если есть, то отправляем его на выход. Если нет, то перекидываем запрос на text-analyzer.
}
