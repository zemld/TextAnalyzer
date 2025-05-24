package handlers

type SaveFileRequest struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}
