package repositories

import "len-test/models"

type StudentRepository interface {
	GetAll() ([]models.Student, error)
	GetById(id string) (models.Student, error)
	Create(studentRequest models.StudentRequest) (models.Student, error)
	Update(studentRequest models.StudentRequest, id string) (models.Student, error)
	Delete(id string) error
}

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id string) (models.Book, error)
	Create(bookRequest models.BookRequest) (models.Book, error)
	Update(bookRequest models.BookRequest, id string) (models.Book, error)
	Delete(id string) error
}

type BorrowingRepository interface {
	GetAll() ([]models.Borrowing, error)
	GetById(id string) (models.Borrowing, error)
	Create(borrowingRequest models.BorrowingRequest) (models.Borrowing, error)
	Update(returningRequest models.ReturningRequest, id string) (models.Borrowing, error)
	Delete(id string) error
}