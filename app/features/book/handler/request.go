package handler

type BookRequest struct {
	Title       string `json:"judul"`
	Description string `json:"tahun"`
	BookImage   string `json:"penerbit"`
}
