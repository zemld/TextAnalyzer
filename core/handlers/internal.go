package handlers

import (
	"crypto/sha256"
	"io"
	"net/http"
)

func getHashOfFileFromRequest(w http.ResponseWriter, r *http.Request) (string, error) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	}
	hash := getHash(content)
	return hash, nil
}

func getHash(fileContent []byte) string {
	hash := sha256.New()
	hash.Write(fileContent)
	return string(hash.Sum(nil))
}
