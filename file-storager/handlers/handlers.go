package handlers

import (
	"net/http"
)

// @description Check if file exists in DB.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} FileExistsResponse
// @failure 401 {object} FileExistsResponse
// @failure 500 {object} FileExistsResponse
// @router /files/exists/{id} [get]
func CheckFileExistsHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Upload file to DB.
// @tag.name File operations
// @param file formData file true "File to upload"
// @param id formData int true "File ID"
// @produce json
// @success 200 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/upload [post]
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Download file from DB.
// @tag.name File operations
// @param id path int true "File ID"
// @produce plain
// @success 200 {file} blob
// @success 401 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/{id} [get]
func GetFileHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Save analysis result to DB.
// @tag.name File operations
// @accept json
// @param id path int true "File ID"
// @produce json
// @success 200 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analysis/{id} [post]
func SaveAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Get analysis result from DB. Result contains amount of paragraphs, sentences, words, symbols. Also contains average amount of sentences per paragraph, words per sentence, length of words.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} AnalysisResponse
// @failure 401 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analysis/{id} [get]
func GetAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Save word cloud to DB.
// @tag.name File operations
// @accept png
// @param id path int true "File ID"
// @param wordCloud formData file true "Word cloud to save"
// @produce json
// @success 200 {object} FileStatusResponse
// @failure 401 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/wordcloud/{id} [post]
func SaveWordCloudHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Get word cloud from DB.
// @tag.name File operations
// @param id path int true "File ID"
// @produce png
// @success 200 {file} blob
// @failure 401 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/wordcloud/{id} [get]
func GetWordCloudHandler(w http.ResponseWriter, r *http.Request) {
}
