package handlers

import (
	"net/http"
)

// @description Uploading file.
// @tag.name File operations
// @param file formData file true "File to upload"
// @produce json
// @success 200 {object} models.FileStatusResponse
// @failure 500 {object} models.FileStatusResponse
// @router /files/upload [post]
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход файл. Считаем хэш от него и перекидываем запрос на file-storager.
	hash, err := getHashOfFileFromRequest(w, r)
	if err != nil {
		return
	}

	client := http.Client{}
	request, _ := http.NewRequest("GET", "http://file-storager-service:8083/files/"+hash, nil)

	_, err = client.Do(request)
}

// @description Downloading file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce plain
// @success 200 {file} blob
// @failure 500 {object} models.FileStatusResponse
// @router /files/download/{id} [get]
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем есть ли файл в базе данных и перекидываем запрос на file-storager.
}

// @description Analyzing file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} models.AnalysisResponse
// @failure 500 {object} models.FileStatusResponse
// @router /files/analyze/{id} [get]
func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем есть ли файл в базе данных.
	// Смотрим, есть ли уже результат анализа. Если есть, то отправляем его на выход. Если нет, то перекидываем запрос на file-analyzer.
}

// @description Getting word cloud.
// @tag.name File operations
// @param id path int true "File ID"
// @produce png
// @success 200 {file} blob
// @failure 500 {object} models.FileStatusResponse
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем есть ли файл в базе данных.
	// Смотрим, есть ли уже результат облако. Если есть, то отправляем его на выход. Если нет, то перекидываем запрос на text-analyzer.
}

// @descriptoin Comparing files.
// @tag.name File operations
// @param first-id path int true "First file ID"
// @param second-id path int true "Second file ID"
// @produce json
// @success 200 {object} models.CompareResponse
// @failure 500 {object} models.FileStatusResponse
func CompareFilesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файлов. Проверяем есть ли файлы в базе данных.
	// Смотрим, есть ли уже результат сравнения. Если есть, то обрабатываем его и возвращаем результат - процент плагиата с указанием что схоже.
	// Если нет результатов анализа хотя бы одного файла, то получаем файл и отдаем его в обработку.
}
