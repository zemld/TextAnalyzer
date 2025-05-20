package handlers

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/zemld/TextAnalyzer/file-storager/db"
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	setAccessControlForOrigin(w, r)
	if err != nil {
		log.Fatalln("Something went wrong while parsing id.")
		writeFileStatusResponse(w, -1, "Wrong id. Cannot upload file.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Parsed id is %d.\n", id)
	if db.CheckFileExistance(id) {
		log.Printf("File with id %d exists in DB.\n", id)
		writeGoodFileExistsResponse(w, id)
		return
	}
	log.Printf("File with id %d does not exist in DB.\n", id)
	writeBadFileExistsResponse(w)
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
	r.ParseMultipartForm(10 << 20)
	setAccessControlForOrigin(w, r)
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		writeFileStatusResponse(w, id, "Wrong id. Cannot upload file.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		writeFileStatusResponse(w, id, "Cannot parse file.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	db.StoreDocument(buf, id, db.FilesCollection)
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
	setAccessControlForOrigin(w, r)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeFileStatusResponse(w, id, "Wrong id. Cannot download file.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := db.GetDocument(id, db.FilesCollection)
	if err != nil {
		writeFileStatusResponse(w, id, "Cannot download file.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	w.Write(file)
}

// TODO: добавить описание json параметра с результатом анализа.
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
