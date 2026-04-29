package models

type Genin struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Category string `json:"category"`
	Aldeia string `json:"aldeia"`
}