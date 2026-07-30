package model

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
