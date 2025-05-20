package handlers

import (
	"io"
	"log"
	"net/http"

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
	setAccessControlForOrigin(w, r)
	id := parseIdFromRequest(w, r)
	if id == -1 {
		return
	}

	log.Printf("Parsed id is %d.\n", id)
	if !db.CheckFileExistance(id) {
		log.Printf("File with id %d does not exist in DB.\n", id)
		writeBadFileExistsResponse(w)
		return
	}
	log.Printf("File with id %d exists in DB.\n", id)
	writeGoodFileExistsResponse(w, id)
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
	setAccessControlForOrigin(w, r)
	id := parseIdFromRequest(w, r)
	if id == -1 {
		return
	}

	r.ParseMultipartForm(10 << 20)
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
	id := parseIdFromRequest(w, r)
	if id == -1 {
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
// @param
// @produce json
// @success 200 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @failure 500 {object} FileStatusResponse
// @router /files/analysis/{id} [post]
func SaveAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
	setAccessControlForOrigin(w, r)
	id := parseIdFromRequest(w, r)
	if id == -1 {
		return
	}

	// TODO: добавить чтение json файла.
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
	setAccessControlForOrigin(w, r)
	id := parseIdFromRequest(w, r)
	if id == -1 {
		return
	}
	// TODO: подумать, как преобразовывать результаты анализа к типу ответа и возвращать его.
	_, _ = db.GetAnalysisResult(id)

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
	setAccessControlForOrigin(w, r)
	id := parseIdFromRequest(w, r)
	if id == -1 {
		return
	}

	r.ParseMultipartForm(10 << 20)
	cloud, _, err := r.FormFile("wordCloud")
	if err != nil {
		writeFileStatusResponse(w, id, "Cannot get word cloud from request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO: пересмотреть работу с картинкой.
	var cloudBytes []byte
	cloudBytes, _ = io.ReadAll(cloud)
	err = db.StoreDocument(cloudBytes, id, db.WordCloudCollection)
	if err != nil {
		writeFileStatusResponse(w, id, "Something went wrong during storing wordcloud.")
		w.WriteHeader(http.StatusInternalServerError)
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
	setAccessControlForOrigin(w, r)
	id := parseIdFromRequest(w, r)
	if id == -1 {
		return
	}

	_, err := db.GetDocument(id, db.WordCloudCollection)
	if err != nil {
		writeFileStatusResponse(w, id, "Something went wrong during getting wordcloud.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: разобраться как возвращать картинку из ручки.
}
