package bookandauthor_entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookID        int
	BookName      string
	BookPageCount int
	BookStock     int
	BookPrice     float64
	BookStockCode string
	BookISBN      string
	BookAuthorID  uint
	Author        Author `gorm:"foreignKey:BookAuthorID;references:ID"`
}
