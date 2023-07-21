package services

import "len-test/repositories"
import "len-test/models"

type BookService struct {
	repository repositories.BookRepository
}

func InitBookService() BookService {
	return BookService{
		repository: &repositories.BookRepositoryImp{},
	}
}

func (us *BookService) Create(bookRequest models.BookRequest) (models.Book, error) {
	return us.repository.Create(bookRequest)
}

func (us *BookService) GetAll() ([]models.Book, error) {
	return us.repository.GetAll()
}

func (us *BookService) GetById(id string) (models.Book, error) {
	return us.repository.GetById(id)
}

func (us *BookService) Update(bookRequest models.BookRequest, id string) (models.Book, error) {
	return us.repository.Update(bookRequest, id)
}

func (us *BookService) Delete(id string) error {
	return us.repository.Delete(id)
}