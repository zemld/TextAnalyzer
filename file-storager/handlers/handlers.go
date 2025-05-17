package handlers

import (
	"net/http"
)

// @description Check if file exists in DB. One of hash or id is required.
// @param id query int true "Id of file"
// @produce json
// @success 200 {json} FileExistsResponse
// @failure 401 {json} FileExistsResponse
// @failure 500 {json} FileExistsResponse
// @router /files/check/{hash} [get]
func CheckFileExistsHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Upload file to DB.
// @param file formData file true "File to upload"
// @param id formData int true "Id of file"
// @produce json
// @success 200 {json} FileUploadDownloadResponse
// @failure 500 {json} FileUploadDownloadResponse

// @router /files/upload [post]
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
}

// @description Download file from DB.
// @param id path int true "Id of file"
// @produce plain
// @success 200 formData file
// @success 401 {json} FileUploadDownloadResponse
// @failure 500 {json} FileUploadDownloadResponse
// @router /files/{id} [get]
func GetFileHandler(w http.ResponseWriter, r *http.Request) {
}

func SaveAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла и результат анализа. Сохраняем результат анализа в бд.
}

func GetAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Берем результат анализа из бд и отправляем его на выход.
}
