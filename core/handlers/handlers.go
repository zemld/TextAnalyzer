package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
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
	id := parseIdFromRequestAndCreateResponse(w, r, getFilePattern)
	if id == -1 {
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusInternalServerError)
		return
	}

	log.Printf("Sending request for downloading file with id: %d", id)
	request, _ := http.NewRequest("GET", "http://file-storager-service:8082/files/"+strconv.Itoa(id), nil)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request for downloading file: %s", err)
		writeFileStatusResponse(w, -1, "Error sending request for downloading file", http.StatusInternalServerError)
		return
	}
	file := response.Body
	defer response.Body.Close()
	content, _ := io.ReadAll(file)
	w.Write(content)
}

// @description Analyzing file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} Analysis
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analyze/{id} [get]
func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	id := parseIdFromRequestAndCreateResponse(w, r, analyzePattern)
	if id == -1 {
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusInternalServerError)
		return
	}
	if !checkFileExistance(w, id) {
		return
	}
	if ok, err := getSavedAnalysis(w, id); ok || err != nil {
		return
	}
	content, ok := getFileFromDB(w, id)
	if !ok {
		return
	}
	analyzeFileAndStoreAnalysisResultIntoDB(w, id, content)
}

// @description Getting word cloud.
// @tag.name File operations
// @param id path int true "File ID"
// @produce png
// @success 200 {file} blob
// @failure 500 {object} FileStatusResponse
// @router /files/wordcloud/{id} [get]
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	id := parseIdFromRequestAndCreateResponse(w, r, wordCloudPattern)
	if id == -1 {
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusInternalServerError)
		return
	}
	if !checkFileExistance(w, id) {
		return
	}
	if ok, err := getSavedWordCloud(w, id); ok || err != nil {
		return
	}
	content, ok := getFileFromDB(w, id)
	if !ok {
		return
	}
	createWordCloudAndStoreItIntoDB(w, id, content)
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
