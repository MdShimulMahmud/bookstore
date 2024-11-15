package types

import validation "github.com/Go-ozzo/ozzo-validation"

type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookname"`
	Author      string `json:"author"`
	Publication string `json:"publication,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.Author,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(5, 50)),
	)
}