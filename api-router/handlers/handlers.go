package handlers

import (
	"io"
	"log"
	"net/http"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("We are at upload file handler")
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "hi from error catcher", http.StatusInternalServerError)
		log.Fatalln(err.Error())
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "There is no file in request", http.StatusBadRequest)
		log.Fatalln(err.Error())
		return
	}

	buffer, contentType, err := createBodyRequest(&file, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err.Error())
		return
	}

	request, _ := http.NewRequest("POST", "http://core:8081/files/upload", buffer)
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
