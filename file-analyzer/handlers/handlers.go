package handlers

import (
	"net/http"
)

// @description Analyze file.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} Analysis
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analyze/{id} [get]
func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {

}

// @description Get word cloud.
// @tag.name File operations
// @param id path int true "File ID"
// @produce png
// @success 200 {object} blob
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
func WordCloudHandler(w http.ResponseWriter, r *http.Request) {

}
