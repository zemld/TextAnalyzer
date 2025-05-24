package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// @description Uploading file.
// @tag.name File operations
// @param file body string true "File to upload"
// @produce json
// @success 200 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/upload [post]
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	content := getFileContent(r)
	log.Printf("Sending request for storing file with content: %s", content)
	request, _ := http.NewRequest("POST", "http://file-storager-service:8082/files/upload", bytes.NewBuffer([]byte(content)))
	request.Header.Set("Content-Type", "text/plain")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request for storing file: %s", err)
		writeFileStatusResponse(w, -1, "Error sending request for storing file", http.StatusInternalServerError)
		return
	}
	var fileStatusResponse FileStatusResponse
	json.NewDecoder(response.Body).Decode(&fileStatusResponse)
	defer response.Body.Close()
	writeFileStatusResponse(w, fileStatusResponse.Id, fileStatusResponse.Status, response.StatusCode)
}

// @description Downloading file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce plain
// @success 200 {file} blob
// @failure 500 {object} FileStatusResponse
// @router /files/download/{id} [get]
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем есть ли файл в базе данных и перекидываем запрос на file-storager.
}

// @description Analyzing file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} Analysis
// @failure 500 {object} FileStatusResponse
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
// @failure 500 {object} FileStatusResponse
// @router /files/wordcloud/{id} [get]
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем есть ли файл в базе данных.
	// Смотрим, есть ли уже результат облако. Если есть, то отправляем его на выход. Если нет, то перекидываем запрос на text-analyzer.
}

// @descriptoin Comparing files.
// @tag.name File operations
// @param first-id query int true "First file ID"
// @param second-id query int true "Second file ID"
// @produce json
// @success 200 {object} CompareResponse
// @failure 500 {object} FileStatusResponse
// @router /files/compare [get]
func CompareFilesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файлов. Проверяем есть ли файлы в базе данных.
	// Смотрим, есть ли уже результат сравнения. Если есть, то обрабатываем его и возвращаем результат - процент плагиата с указанием что схоже.
	// Если нет результатов анализа хотя бы одного файла, то получаем файл и отдаем его в обработку.
}
