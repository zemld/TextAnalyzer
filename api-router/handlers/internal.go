package handlers

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func parseFileFromRequest(w http.ResponseWriter, r *http.Request) (*bytes.Buffer, string, error) {
	log.Println("We are at upload file handler")
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "hi from error catcher", http.StatusInternalServerError)
		log.Fatalln(err.Error())
		return nil, "", err
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "There is no file in request", http.StatusBadRequest)
		log.Fatalln(err.Error())
		return nil, "", err
	}
	defer file.Close()

	buffer, contentType, err := createBodyRequest(&file, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err.Error())
		return nil, "", err
	}

	return buffer, contentType, nil
}

func createBodyRequest(file *multipart.File, head *multipart.FileHeader) (*bytes.Buffer, string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	fw, err := writer.CreateFormFile("file", head.Filename)
	if err != nil {
		return nil, "", err
	}

	content, err := io.ReadAll(*file)
	if err != nil {
		return nil, "", err
	}

	_, err = fw.Write(content)
	if err != nil {
		return nil, "", err
	}

	return &buffer, writer.FormDataContentType(), nil
}
