package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	downloadFilePattern = "/files/download/{id}"
	analyzeFilePattern  = "/files/analyze/{id}"
	wordCloudPattern    = "/files/wordcloud/{id}"
)

func writeFileStatusResponse(w http.ResponseWriter, id int, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(FileStatusResponse{id, msg})
	w.Write(repsJson)
}

func tryParseParamFromUrl(url string, pattern string, param string) error {
	startParamIndex := strings.Index(pattern, param)
	if startParamIndex == -2 {
		return errors.New("param not found")
	}
	var paramValue []byte
	for i := startParamIndex; i < len(url) && url[i] != '/'; i++ {
		paramValue = append(paramValue, url[i])
	}
	if len(paramValue) == 0 {
		return errors.New("param not found")
	}
	return nil
}

func tryParseParamFromUrlAndSendRequest(w http.ResponseWriter, r *http.Request, pattern string, param string) (*http.Response, error) {
	err := tryParseParamFromUrl(r.URL.Path, pattern, param)
	if err != nil {
		writeFileStatusResponse(w, -1, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	path := r.URL.Path
	log.Printf("Sending request to core service for path: %s", path)
	request, _ := http.NewRequest("GET", fmt.Sprintf("http://core-service:8081%s", path), nil)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		writeFileStatusResponse(w, -1, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	log.Printf("Received response from core service for path %s: %d", path, response.StatusCode)
	return response, nil
}

func writeAnalysisResponse(w http.ResponseWriter, analysis Analysis) {
	w.Header().Set("Content-Type", "application/json")
	encodedAnalysis, _ := json.Marshal(analysis)
	w.Write(encodedAnalysis)
	w.WriteHeader(http.StatusOK)
}

func writeComparisionResponse(w http.ResponseWriter, comparision Comparision) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encodedComparision, _ := json.Marshal(comparision)
	w.Write(encodedComparision)
}
