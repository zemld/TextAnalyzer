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

func CompareFilesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файлов. Проверяем есть ли файлы в базе данных.
	// Смотрим, есть ли уже результат сравнения. Если есть, то обрабатываем его и возвращаем результат - процент плагиата с указанием что схоже.
	// Если нет результатов анализа хотя бы одного файла, то получаем файл и отдаем его в обработку.
}
