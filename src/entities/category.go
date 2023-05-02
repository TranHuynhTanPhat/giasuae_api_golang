package entities

type Category struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       int    `json:"type"`
	Created_at string `json:"created_at"`
}
