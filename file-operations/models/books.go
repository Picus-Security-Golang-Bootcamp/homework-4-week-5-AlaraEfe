package models

type Book struct {
	BookID         int     `json:"BookID"`
	BookName       string  `json:"BookName"`
	BookPageCount  int     `json:"BookPageCount"`
	BookStock      int     `json:"BookStock"`
	BookPrice      float64 `json:"BookPrice"`
	BookStockCode  string  `json:"BookStockCode"`
	BookISBN       string  `json:"BookISBN"`
	BookAuthorID   int     `json:"BookAuthorID"`
	BookAuthorName string  `json:"BookAuthorName"`
}

type BooksSlice []Book
