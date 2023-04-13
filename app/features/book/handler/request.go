package handler

type BookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	BookImage   string `json:"book_image"`
}
