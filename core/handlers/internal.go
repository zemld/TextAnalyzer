package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
