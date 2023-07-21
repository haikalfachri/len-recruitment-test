package repositories

import (
	"len-test/database"
	"len-test/models"
)

type BookRepositoryImp struct {
}

func InitBookRepository() BookRepository {
	return &BookRepositoryImp{}
}

func (ur *BookRepositoryImp) Create(bookRequest models.BookRequest) (models.Book, error) {
	var book models.Book = models.Book{
		Title: bookRequest.Title,
		Author: bookRequest.Author,
		Quantity: bookRequest.Quantity,
		Storage: bookRequest.Storage,
	}

	if err := database.DB.Create(&book).Error; err != nil {
		return models.Book{}, err
	}

	if err := database.DB.Last(&book).Error; err != nil {
		return models.Book{}, err
	}

    return book, nil
}

func (ur *BookRepositoryImp) GetAll() ([]models.Book, error) {
	var books []models.Book

	if err := database.DB.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (ur *BookRepositoryImp) GetById(id string) (models.Book, error) {
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (ur *BookRepositoryImp) Update(bookRequest models.BookRequest, id string) (models.Book, error) {
	book, err := ur.GetById(id)

	if bookRequest.Title != book.Title {
		book.Title = bookRequest.Title
	}
	if bookRequest.Author != book.Author{
		book.Author = bookRequest.Author
	}
	if bookRequest.Quantity != book.Quantity{
		book.Quantity = bookRequest.Quantity
	}
	if bookRequest.Storage != book.Storage{
		book.Storage = bookRequest.Storage
	}
	
	if err != nil {
		return models.Book{}, err
	}

	if err := database.DB.Save(&book).Error; err != nil {
		return models.Book{}, err
	}

    return book, nil
}

func (ur *BookRepositoryImp) Delete(id string) error {
	book, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		return err
	}

    return nil
}

