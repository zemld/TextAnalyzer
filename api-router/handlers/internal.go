package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
)

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
