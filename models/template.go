package models

type Template struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
	HTML    bool   `json:"html"`
}
