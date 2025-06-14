package server

type CreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
