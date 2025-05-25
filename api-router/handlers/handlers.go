package handlers

import (
	"bytes"
	"encoding/json"
	"io"
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
	body := r.Body
	defer body.Close()
	fileContent, err := io.ReadAll(body)
	if err != nil {
		writeFileStatusResponse(w, -1, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(fileContent) == 0 {
		writeFileStatusResponse(w, -1, "File is empty", http.StatusBadRequest)
		return
	}
	contentToSend := bytes.NewBuffer(fileContent)
	request, _ := http.NewRequest("POST", "http://core-service:8081/files/upload", contentToSend)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		writeFileStatusResponse(w, -1, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	var fileStatus FileStatusResponse
	json.Unmarshal(responseBody, &fileStatus)
	writeFileStatusResponse(w, fileStatus.Id, fileStatus.Status, response.StatusCode)
}

// @description Downloading file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce plain
// @produce json
// @success 200 {file} blob
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/download/{id} [get]
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	response, err := tryParseParamFromUrlAndSendRequest(w, r, downloadFilePattern, "{id}")
	if err != nil {
		return
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	status := response.StatusCode
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(responseBody)
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
	response, err := tryParseParamFromUrlAndSendRequest(w, r, analyzeFilePattern, "{id}")
	if err != nil {
		return
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	status := response.StatusCode
	var analysis Analysis
	if err = json.Unmarshal(responseBody, &analysis); err != nil {
		writeFileStatusResponse(w, -1, "Cannot get result.", status)
		return
	}
	writeAnalysisResponse(w, analysis)
}

// @description Getting word cloud.
// @tag.name File operations
// @param id path int true "File ID"
// @produce png
// @produce json
// @success 200 {png} blob
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/wordcloud/{id} [get]
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	response, err := tryParseParamFromUrlAndSendRequest(w, r, wordCloudPattern, "{id}")
	if err != nil {
		return
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	status := response.StatusCode
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "image/png")
	w.Write(responseBody)
}

// @description Comparing files.
// @tag.name File operations
// @param first-id query int true "First file ID"
// @param second-id query int true "Second file ID"
// @produce json
// @success 200 {object} Comparision
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/compare [get]
func CompareFilesHandler(w http.ResponseWriter, r *http.Request) {
	firstId, err := strconv.Atoi(r.URL.Query().Get("first-id"))
	if err != nil {
		writeFileStatusResponse(w, -1, "Cannot get first file ID.", http.StatusBadRequest)
		return
	}
	secondId, err := strconv.Atoi(r.URL.Query().Get("second-id"))
	if err != nil {
		writeFileStatusResponse(w, -1, "Cannot get second file ID.", http.StatusBadRequest)
		return
	}
	request, _ := http.NewRequest("GET", "http://core-service:8081/files/compare", nil)
	q := request.URL.Query()
	q.Add("first-id", strconv.Itoa(firstId))
	q.Add("second-id", strconv.Itoa(secondId))
	request.URL.RawQuery = q.Encode()
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		writeFileStatusResponse(w, -1, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	var comparision Comparision
	if err = json.Unmarshal(responseBody, &comparision); err != nil {
		writeFileStatusResponse(w, -1, "Cannot get result.", response.StatusCode)
	}
	writeComparisionResponse(w, comparision)
}
