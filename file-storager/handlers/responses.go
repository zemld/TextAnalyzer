package handlers

type FileExistsResponse struct {
	Exists bool   `json:"exists"`
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type FileStatusResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}
