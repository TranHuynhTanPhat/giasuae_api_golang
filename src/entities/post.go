package entities

type Post struct {
	ID         int    `json:"id"`
	Type       int    `json:"type"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Img        string `json:"img"`
	Created_at string `json:"created_at"`
}
