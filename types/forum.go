package types

type Forum struct {
	Title  string `json:"title"`
	Writer string `json:"writer"`
	Detail string `json:"detail"`
}

type ForumResponse struct {
	*ApiResponse
	*Forum
}
