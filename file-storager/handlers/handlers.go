package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// @description Check if file exists in DB.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} FileExistsResponse
// @failure 400 {object} FileStatusResponse
// @failure 500 {object} FileExistsResponse
// @router /files/exists/{id} [get]
func CheckFileExistsHandler(w http.ResponseWriter, r *http.Request) {
	setAccessControlForOrigin(w, r)
	id := parseIdFromRequestAndCreateResponse(w, r, existsPattern)
	if id == -1 {
		return
	}

	log.Printf("Parsed id is %d.\n", id)
	if !checkFileExistance(id) {
		log.Printf("File with id %d does not exist in DB.\n", id)
		writeBadFileExistsResponse(w)
		return
	}
	log.Printf("File with id %d exists in DB.\n", id)
	writeGoodFileExistsResponse(w, id)
}

// @description Upload file to DB.
// @tag.name File operations
// @param file body string true "File to upload"
// @produce json
// @success 200 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/upload [post]
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	setAccessControlForOrigin(w, r)
	text, ok := parseSaveFileRequestAndCreateResponse(w, r)
	if !ok {
		return
	}
	log.Print(text)
	hash := getHash(text)
	id := storeHash(hash)
	if id == -1 {
		writeFileStatusResponse(w, -1, "Cannot store file.",
			http.StatusInternalServerError)
		return
	}
	buf := []byte(text)
	if err := storeDocument(buf, id); err != nil {
		writeFileStatusResponse(w, id, "Cannot store file.",
			http.StatusInternalServerError)
		return
	}
	writeFileStatusResponse(w, id, "File uploaded.", http.StatusOK)
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
	id := parseIdFromRequestAndCreateResponse(w, r, getPattern)
	if id == -1 {
		return
	}

	file, err := getDocument(id)
	if err != nil {
		writeFileStatusResponse(w, id, "Cannot download file.",
			http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	w.Write(file)
	w.WriteHeader(http.StatusOK)
}

// @description Save analysis result to DB.
// @tag.name File operations
// @accept json
// @param analysis body Analysis true "Result of file analysis"
// @produce json
// @success 200 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analysis [post]
func SaveAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	analysis := Analysis{}
	err := decoder.Decode(&analysis)
	if err != nil {
		writeFileStatusResponse(w, -1, "Cannot read analysis result.",
			http.StatusInternalServerError)
		return
	}
	err = storeAnalysisResult(analysis)
	if err != nil {
		writeFileStatusResponse(w, -1, "Cannot save analysis result.",
			http.StatusInternalServerError)
		return
	}
	writeFileStatusResponse(w, analysis.Id, "Analysis result saved.", http.StatusOK)
}

// @description Get analysis result from DB. Result contains amount of paragraphs, sentences, words, symbols. Also contains average amount of sentences per paragraph, words per sentence, length of words.
// @tag.name File operations
// @param id path int true "File ID"
// @produce json
// @success 200 {object} Analysis
// @failure 401 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analysis/{id} [get]
func GetAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
	id := parseIdFromRequestAndCreateResponse(w, r, analysisPattern)
	if id == -1 {
		return
	}
	mappedAnalysisResult, _ := getAnalysisResult(id)
	writeAnalysisResponse(w,
		makeResponseFromSelectingAnalysisResult(id, mappedAnalysisResult))
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
	id := parseIdFromRequestAndCreateResponse(w, r, wordCloudPattern)
	if id == -1 {
		return
	}

	cloud, _, err := r.FormFile("wordCloud")
	if err != nil {
		writeFileStatusResponse(w, id,
			"Cannot get word cloud from request.",
			http.StatusBadRequest)
		return
	}

	if err = storeWordCloud(id, cloud); err != nil {
		writeFileStatusResponse(w, id,
			"Something went wrong during storing wordcloud.",
			http.StatusInternalServerError)
		return
	}

	writeGoodFileExistsResponse(w, id)
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
	id := parseIdFromRequestAndCreateResponse(w, r, wordCloudPattern)
	if id == -1 {
		return
	}

	cloud, err := getWordCloud(id)
	if err != nil {
		writeFileStatusResponse(w, id,
			"Something went wrong during getting wordcloud.",
			http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(cloud)
}
