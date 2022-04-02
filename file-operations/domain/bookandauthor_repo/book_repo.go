package bookandauthor_repo

import (
	"errors"
	"fmt"
	"github.com/AlaraEfe/file-operations/file-operations/domain/bookandauthor_entities"
	"github.com/AlaraEfe/file-operations/file-operations/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&bookandauthor_entities.Book{})
}

func (b *BookRepository) InsertBookData(booksModel models.BooksSlice) {

	books := []bookandauthor_entities.Book{}

	for _, book := range booksModel {

		books = append(books, bookandauthor_entities.Book{
			BookID:        book.BookID,
			BookName:      book.BookName,
			BookPageCount: book.BookPageCount,
			BookStock:     book.BookStock,
			BookPrice:     book.BookPrice,
			BookStockCode: book.BookStockCode,
			BookISBN:      book.BookISBN,
			BookAuthorID:  uint(book.BookAuthorID),
		})

	}

	for _, book := range books {
		b.db.Unscoped().Where(bookandauthor_entities.Book{BookID: book.BookID}).Attrs(bookandauthor_entities.Book{BookID: book.BookID, BookName: book.BookName}).FirstOrCreate(&book)

	}

}

func (b *BookRepository) ListAllBooksWithRawSQL() []bookandauthor_entities.Book {
	var books []bookandauthor_entities.Book
	b.db.Raw("SELECT * FROM books").Scan(&books)

	return books
}

func (b *BookRepository) SearchByNameWithRawSQL(name string) []bookandauthor_entities.Book {
	var books []bookandauthor_entities.Book
	b.db.Raw("SELECT * FROM books WHERE book_name ILIKE ?", "%"+name+"%").Find(&books)

	return books
}

func (b *BookRepository) SoftDeleteGormSQL(ID int) *bookandauthor_entities.Book {
	var book bookandauthor_entities.Book
	b.db.Where("book_id = ?", ID).Delete(&book)

	return &book
}

func (b *BookRepository) BuyBookWithRawSQL(ID int, quantitiy int) (*bookandauthor_entities.Book, error) {
	var book bookandauthor_entities.Book

	result := b.db.Raw("SELECT * FROM books WHERE book_id = ?", ID).First(&book)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error

	}

	if book.BookStock >= quantitiy {

		b.db.Model(&book).UpdateColumn("book_stock", gorm.Expr("book_stock - ?", quantitiy))

	} else {
		fmt.Println("There is not enough stock of that book")
		return nil, result.Error
	}

	bookPrice := book.BookPrice

	totalBookCost := bookPrice * float64(quantitiy)

	s := fmt.Sprintf("The total cost of %d book/books named '%s' is: %2.f\n ", quantitiy, book.BookName, totalBookCost)
	print(s)

	return &book, nil
}

func (b *BookRepository) GetByID(id int) (*bookandauthor_entities.Book, error) {
	var book bookandauthor_entities.Book

	result := b.db.Where("book_id = ?", id).First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		softDeletedRecords := b.db.Unscoped().Where("book_id = ?", id).First(&book)
		if errors.Is(softDeletedRecords.Error, gorm.ErrRecordNotFound) {

			return nil, result.Error

		}

		fmt.Println("That book was deleted from Book Archive database")

		return &book, nil

	}

	return &book, nil
}

func (b *BookRepository) FindByName(bookname string) (*bookandauthor_entities.Book, error) {
	var book bookandauthor_entities.Book
	result := b.db.Where("book_name ILIKE ?", "%"+bookname+"%").First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error

	}

	return &book, nil
}

func (b *BookRepository) GetAuthorWithBooks(bookName string) (*bookandauthor_entities.Author, error) {
	var book bookandauthor_entities.Book
	var author bookandauthor_entities.Author
	result := b.db.Where("book_name ILIKE ?", "%"+bookName+"%").First(&book)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error

	}

	resultAuthor := b.db.Preload("authors").Where("ID = ?", book.BookAuthorID).First(&author)

	if errors.Is(resultAuthor.Error, gorm.ErrRecordNotFound) {
		return nil, resultAuthor.Error

	}

	s := fmt.Sprintf("The author of %s is %s.\n ", book.BookName, author.AuthorName)
	print(s)

	return &author, nil
}
