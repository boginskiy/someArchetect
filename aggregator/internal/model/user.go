package model

type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}
