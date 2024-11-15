package services

import (
	"errors"
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
)

type BookService struct {
    repo domain.IBookRepo
}

func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
    return &BookService{
        repo: bookRepo,
    }
}

func (service *BookService) GetBooks(bookID uint) ([]types.BookRequest, error) {
    var allBooks []types.BookRequest
    book := service.repo.GetBooks(bookID)
    if len(book) == 0 {
        return nil, errors.New("no book found")
    }
    for _, val := range book {
        allBooks = append(allBooks, types.BookRequest{
            ID:          val.ID,
            BookName:    val.BookName,
            Author:      val.Author,
            Publication: val.Publication,
        })
    }
    return allBooks, nil
}

func (service *BookService) CreateBook(book *models.Book) error {
    if err := service.repo.CreateBook(book); err != nil {
        return errors.New("book was not created")
    }
    return nil
}
func (service *BookService) UpdateBook(book *models.Book) error {
    if err := service.repo.UpdateBook(book); err != nil {
        return errors.New("book update was unsuccesful")
    }
    return nil
}



func (service *BookService) DeleteBook(bookID uint) error {
    if err := service.repo.DeleteBook(bookID); err != nil {
        return errors.New("book deletion was unsuccessful")
    }
    return nil
}