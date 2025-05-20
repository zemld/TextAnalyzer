package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

const (
	incorrectIdMsg = "Incorrect ID."
)

// TODO: имеет смысл добавить кастомный тип для id.

func parseIdFromRequest(w http.ResponseWriter, r *http.Request) int {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeFileStatusResponse(w, -1, incorrectIdMsg, http.StatusBadRequest)
		return -1
	}
	return id
}

func writeFileExistsResponse(w http.ResponseWriter, resp FileExistsResponse) {
	w.Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(resp)
	w.Write(repsJson)
}

func writeBadFileExistsResponse(w http.ResponseWriter) {
	resp := FileExistsResponse{Exists: false, Id: -1, Status: "File doesn't exist"}
	writeFileExistsResponse(w, resp)
}

func writeGoodFileExistsResponse(w http.ResponseWriter, id int) {
	resp := FileExistsResponse{Exists: true, Id: id, Status: "File exists"}
	writeFileExistsResponse(w, resp)
}

func writeFileStatusResponse(w http.ResponseWriter, id int, msg string, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(FileStatusResponse{id, msg})
	w.Write(repsJson)
	w.WriteHeader(statusCode)
}

func setAccessControlForOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func makeResponseFromSelectingAnalysisResult(id int, result map[string]any) Analysis {
	// TODO: вспомнить, как пишется type assertion.
	return Analysis{Id: id}
}
