package handlers

import (
	"io"
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
	buffer, contentType, err := parseFileFromRequest(w, r)
	if err != nil {
		return
	}

	request, _ := http.NewRequest("POST", "http://core-service:8081/files/upload", buffer)
	request.Header.Set("Content-Type", contentType)
	c := http.Client{}
	response, err := c.Do(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(response.StatusCode)
	responseBody, _ := io.ReadAll(response.Body)
	w.Write(responseBody)
}

// @description Downloading file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce plain
// @produce json
// @success 200 {file} blob
// @failure 500 {object} models.FileStatusResponse
// @router /files/download/{id} [get]
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Прокидываем запрос дальше на ядро.
}

// @description Analyzing file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} models.AnalysisResponse
// @failure 500 {object} models.FileStatusResponse
// @router /files/analyze/{id} [get]
func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Прокидываем запрос дальше на ядро.
}

// @description Getting word cloud.
// @tag.name File operations
// @param id path int true "File ID"
// @produce png
// @produce json
// @success 200 {png} blob
// @failure 500 {object} models.FileStatusResponse
// @router /files/wordcloud/{id} [get]
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Прокидываем запрос дальше на ядро.
}

// @description Comparing files.
// @tag.name File operations
// @param first-id path int true "First file ID"
// @param second-id path int true "Second file ID"
// @produce json
// @success 200 {object} models.CompareResponse
// @failure 500 {object} models.FileStatusResponse
// @router /files/compare/{first-id}/{second-id} [get]
func CompareFilesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файлов. Прокидываем запрос дальше на ядро.
}
