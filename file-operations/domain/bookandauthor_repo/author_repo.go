package bookandauthor_repo

import (
	"github.com/AlaraEfe/file-operations/file-operations/domain/bookandauthor_entities"
	"github.com/AlaraEfe/file-operations/file-operations/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (b *AuthorRepository) Migrations() {
	b.db.AutoMigrate(&bookandauthor_entities.Author{})
}

func (b *AuthorRepository) InsertAuthorData(authorsModel models.BooksSlice) {

	authors := []bookandauthor_entities.Author{}
	for _, author := range authorsModel {
		authors = append(authors, bookandauthor_entities.Author{
			AuthorName: author.BookAuthorName,
			ID:         uint(author.BookAuthorID),
		})

	}

	for _, author := range authors {
		b.db.Where(bookandauthor_entities.Author{ID: author.ID}).Attrs(bookandauthor_entities.Author{AuthorName: author.AuthorName}).FirstOrCreate(&author)

	}

}

func (a *AuthorRepository) GetBooksWithAuthorName(authorName string) []bookandauthor_entities.Book {
	var author bookandauthor_entities.Author
	var books []bookandauthor_entities.Book

	a.db.Where("author_name ILIKE ?", "%"+authorName+"%").Find(&author)

	a.db.Preload("books").Where("book_author_id = ?", author.ID).Find(&books)

	return books
}
