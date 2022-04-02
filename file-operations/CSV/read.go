package csv_utils

import (
	"encoding/csv"
	"github.com/AlaraEfe/file-operations/file-operations/models"
	"os"
	"strconv"
)

func ReadCSV(filename string) (models.BooksSlice, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books models.BooksSlice

	for _, line := range records[1:] {

		book := models.Book{}
		book.BookID, _ = strconv.Atoi(line[0])
		book.BookName = line[1]
		book.BookPageCount, _ = strconv.Atoi(line[2])
		book.BookStock, _ = strconv.Atoi(line[3])
		book.BookPrice, _ = strconv.ParseFloat(line[4], 64)
		book.BookStockCode = line[5]
		book.BookISBN = line[6]
		book.BookAuthorID, _ = strconv.Atoi(line[7])
		book.BookAuthorName = line[8]

		books = append(books, book)

	}

	return books, nil
}
