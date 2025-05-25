package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	downloadFilePattern = "/files/download/{id}"
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
