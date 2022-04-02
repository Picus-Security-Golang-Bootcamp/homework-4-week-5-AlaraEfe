package bookandauthor_entities

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	AuthorName string `gorm:"unique"`
	Books      []Book `gorm:"foreignKey:BookAuthorID"`
}
