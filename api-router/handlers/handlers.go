package handlers

import (
	"io"
	"net/http"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	buffer, contentType, err := parseFileFromRequest(w, r)
	if err != nil {
		return
	}

	request, _ := http.NewRequest("POST", "http://core-service:8081/files/upload", buffer)
	request.Header.Set("Content-Type", contentType)
	c := http.Client{}
	response, err := c.Do(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(response.StatusCode)
	responseBody, _ := io.ReadAll(response.Body)
	w.Write(responseBody)
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Прокидываем запрос дальше на ядро.
}

func AnalyzeFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Прокидываем запрос дальше на ядро.
}

func WordCloudHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Прокидываем запрос дальше на ядро.
}

func CompareFilesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файлов. Прокидываем запрос дальше на ядро.
}
