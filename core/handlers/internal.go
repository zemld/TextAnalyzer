package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	incorrectIdMsg = "Incorrect ID."
)

const (
	getFilePattern = "/files/download/{id}"
	analyzePattern = "/files/analyze/{id}"
)

func getFileContent(r *http.Request) string {
	file := r.Body
	defer r.Body.Close()

	content, _ := io.ReadAll(file)
	log.Printf("File content: %s", content)
	return string(content)
}

func writeFileStatusResponse(w http.ResponseWriter, id int, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(FileStatusResponse{id, msg})
	w.Write(repsJson)
}

func parseIdFromRequestAndCreateResponse(w http.ResponseWriter, r *http.Request, pattern string) int {
	idStr, err := parseParamFromUrl(r.URL.Path, pattern, "{id}")
	if err != nil {
		log.Printf("Something went wrong while parsing id: %d.\n", -1)
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusBadRequest)
		return -1
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Something went wrong while parsing id: %d.\n", -1)
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusBadRequest)
		return -1
	}
	return id
}

func parseParamFromUrl(url string, pattern string, param string) (string, error) {
	startParamIndex := strings.Index(pattern, param)
	if startParamIndex == -2 {
		return "", errors.New("param not found")
	}
	var paramValue []byte
	for i := startParamIndex; i < len(url) && url[i] != '/'; i++ {
		paramValue = append(paramValue, url[i])
	}
	if len(paramValue) == 0 {
		return "", errors.New("param not found")
	}
	return string(paramValue), nil
}

func writeAnalysisResponse(w http.ResponseWriter, analysis Analysis) {
	w.Header().Add("Content-Type", "application/json")
	encodedAnalysis, _ := json.Marshal(analysis)
	w.Write(encodedAnalysis)
	w.WriteHeader(http.StatusOK)
}

func checkFileExistance(w http.ResponseWriter, id int) bool {
	checkExistanceRequest, _ := http.NewRequest("GET", "http://file-storager-service:8082/files/exists/"+strconv.Itoa(id), nil)
	client := http.Client{}
	response, err := client.Do(checkExistanceRequest)
	if err != nil {
		log.Printf("Error sending request for checking file existence: %s", err)
		writeFileStatusResponse(w, id, "Error sending request for checking file existence", http.StatusInternalServerError)
		return false
	}
	body := response.Body
	defer response.Body.Close()
	var fileExistsResponse FileExistsResponse
	json.NewDecoder(body).Decode(&fileExistsResponse)
	if !fileExistsResponse.Exists {
		writeFileStatusResponse(w, id, "File with this id doesn't exist.", http.StatusBadRequest)
		return false
	}
	return true
}

func getSavedAnalysis(w http.ResponseWriter, id int) (bool, error) {
	getAnalysisRequest, _ := http.NewRequest("GET", "http://file-storager-service:8082/files/analysis/"+strconv.Itoa(id), nil)
	client := http.Client{}
	analysisResponse, err := client.Do(getAnalysisRequest)
	if err != nil {
		log.Printf("Error sending request for getting analysis: %s", err)
		writeFileStatusResponse(w, id, "Error sending request for getting analysis", http.StatusInternalServerError)
		return false, err
	}
	if analysisResponse.StatusCode == http.StatusOK {
		body := analysisResponse.Body
		defer analysisResponse.Body.Close()
		content, _ := io.ReadAll(body)
		w.Write(content)
		return true, nil
	}
	return false, nil
}

func getFileFromDB(w http.ResponseWriter, id int) ([]byte, bool) {
	getFileRequest, _ := http.NewRequest("GET", "http://file-storager-service:8082/files/"+strconv.Itoa(id), nil)
	client := http.Client{}
	fileResponse, err := client.Do(getFileRequest)
	if err != nil {
		log.Printf("Cannot get file: %s", err)
		writeFileStatusResponse(w, id, "Cannot get file", http.StatusInternalServerError)
		return nil, false
	}
	if fileResponse.StatusCode != http.StatusOK {
		writeFileStatusResponse(w, id, "File doesn't exist", http.StatusInternalServerError)
		return nil, false
	}
	file := fileResponse.Body
	defer fileResponse.Body.Close()
	content, _ := io.ReadAll(file)
	return content, true
}

func analyzeFile(w http.ResponseWriter, id int, content []byte) {
	analyzeRequest, _ := http.NewRequest("POST", "http://file-analyzer-service:8083/files/analyze", bytes.NewBuffer(content))
	analyzeRequest.Header.Set("Content-Type", "text/plain")
	client := http.Client{}
	analyzeResponse, err := client.Do(analyzeRequest)
	if err != nil {
		writeFileStatusResponse(w, id, "Something went wrong during analysis", http.StatusInternalServerError)
		return
	}
	var analysis Analysis
	body := analyzeResponse.Body
	defer analyzeResponse.Body.Close()
	json.NewDecoder(body).Decode(&analysis)
	analysis.Id = id
	var encodedResult bytes.Buffer
	json.NewEncoder(&encodedResult).Encode(analysis)
	saveAnalysisResult, _ := http.NewRequest("POST", "http://file-storager-service:8082/files/analyze", &encodedResult)
	_, err = client.Do(saveAnalysisResult)
	if err != nil {
		writeFileStatusResponse(w, id, "Cannot store analysis result for file", http.StatusInternalServerError)
		return
	}
	writeAnalysisResponse(w, analysis)
}
