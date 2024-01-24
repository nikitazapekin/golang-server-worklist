package model

// BookList - list of book
type BookList struct {
	Books []Book `json:"books"`
}
 
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"title"`
	Author string `json:"author"`
}