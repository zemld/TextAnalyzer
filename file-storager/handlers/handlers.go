package handlers

import (
	"net/http"
)

// @description Check if file exists in DB.
// @tag.name File operations
// @param id path int true "Id of file"
// @produce json
// @success 200 body FileExistsResponse
// @failure 401 body FileExistsResponse
// @failure 500 body FileExistsResponse
// @router /files/check/{id} [get]
func CheckFileExistsHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Upload file to DB.
// @tag.name File operations
// @accept plain
// @param file formData file true "File to upload"
// @param id formData int true "Id of file"
// @produce json
// @success 200 body FileStatusResponse
// @failure 500 body FileStatusResponse
// @router /files/upload [post]
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Download file from DB.
// @tag.name File operations
// @param id path int true "Id of file"
// @produce plain
// @success 200 formData file
// @success 401 body FileStatusResponse
// @failure 500 body FileStatusResponse
// @router /files/{id} [get]
func GetFileHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Save analysis result to DB.
// @tag.name File operations
// @accept json
// @param id path int true "Id of file"
// @produce plain
// @success 200 body FileStatusResponse
// @failure 500 body FileStatusResponse
// @router /files/analysis/{id} [post]
func SaveAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Get analysis result from DB. Result contains amount of paragraphs, sentences, words, symbols. Also contains average amount of sentences per paragraph, words per sentence, length of words.
// @tag.name File operations
// @param id path int true "Id of file"
// @produce json
// @success 200 body AnalysisResponse
// @failure 500 body FileStatusResponse
func GetAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
}
