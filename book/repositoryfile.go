package book

import (
	"fmt"
)

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book

	fmt.Println("Find all")

	return books, nil
}

func (r *fileRepository) FindByID(ID int) (Book, error) {
	var book Book
	fmt.Println("Find By ID")
	return book, nil
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("Create")
	return book, nil
}
