package handlers

import (
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
