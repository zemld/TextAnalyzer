package handlers

import (
	"net/http"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход файл. Берем хэш от него и обращаемся в бд.
	// если есть, то возвращаем айди файла. Если нет, то сохраняем в бд и возвращаем сгенерированное айди файла.
}

func GetFileHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Проверяем наличие файла и берем файл из бд и отправляем его на выход.
}

func SaveAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла и результат анализа. Сохраняем результат анализа в бд.
}

func GetAnalysisResultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: получаем на вход айди файла. Берем результат анализа из бд и отправляем его на выход.
}
