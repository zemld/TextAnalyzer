package handlers

import (
	"encoding/json"
	"net/http"
)

func writeFileExistsResponse(w *http.ResponseWriter, resp FileExistsResponse) {
	(*w).Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(resp)
	(*w).Write(repsJson)
}

func writeBadFileExistsResponse(w *http.ResponseWriter) {
	resp := FileExistsResponse{Exists: false, Id: -1, Status: "File doesn't exist"}
	writeFileExistsResponse(w, resp)
}

func writeGoodFileExistsResponse(w *http.ResponseWriter, id int) {
	resp := FileExistsResponse{Exists: true, Id: id, Status: "File exists"}
	writeFileExistsResponse(w, resp)
}

func writeFileStatusResponse(w *http.ResponseWriter, id int, msg string) {
	(*w).Header().Add("Content-Type", "application/json")
	repsJson, _ := json.Marshal(FileStatusResponse{id, msg})
	(*w).Write(repsJson)
}
