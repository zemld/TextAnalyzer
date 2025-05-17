package handlers

type FileExistsResponse struct {
	Exists bool   `json:"exists"`
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type FileUploadDownloadResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}
