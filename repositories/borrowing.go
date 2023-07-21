package repositories

import (
	"errors"
	"len-test/database"
	"len-test/models"
	"time"
)

type BorrowingRepositoryImp struct {
}

func InitBorrowingRepository() BorrowingRepository {
	return &BorrowingRepositoryImp{}
}

func (ur *BorrowingRepositoryImp) Create(borrowingRequest models.BorrowingRequest) (models.Borrowing, error) {
	var borrowing models.Borrowing = models.Borrowing{
		StudentID:    borrowingRequest.StudentID,
		BookID:       borrowingRequest.BookID,
		BorrowedAt:   time.Now(),
		BorrowedTerm: borrowingRequest.BorrowedTerm,
	}

	var countBorrowedBook int
	var borrowedBooks []models.Borrowing

	if err := database.DB.Find(&borrowedBooks, `student_id = ?`, borrowingRequest.StudentID).Error; err != nil {
		return models.Borrowing{}, err
	}

	for _, book := range borrowedBooks {
		if book.ReturnedAt.IsZero() {
			countBorrowedBook++
			if countBorrowedBook > 10 {
				err := errors.New("You have reached the limit of borrowing books")
				return models.Borrowing{}, err
			}
		}
	}

	var borrowedBook models.Book

	if err := database.DB.First(&borrowedBook, "id = ?", borrowingRequest.BookID).Error; err != nil {
		return models.Borrowing{}, err
	}

	if borrowedBook.Quantity-1 < 0 {
		err := errors.New("Books not available")
		return models.Borrowing{}, err
	}

	borrowedBook.Quantity -= 1

	if err := database.DB.Create(&borrowing).Error; err != nil {
		return models.Borrowing{}, err
	}

	if err := database.DB.Last(&borrowing).Error; err != nil {
		return models.Borrowing{}, err
	}

	if err := database.DB.Save(&borrowedBook).Error; err != nil {
		return models.Borrowing{}, err
	}

	return borrowing, nil
}

func (ur *BorrowingRepositoryImp) GetAll() ([]models.Borrowing, error) {
	var borrowings []models.Borrowing

	if err := database.DB.Find(&borrowings).Error; err != nil {
		return borrowings, err
	}
	return borrowings, nil
}

func (ur *BorrowingRepositoryImp) GetById(id string) (models.Borrowing, error) {
	var borrowing models.Borrowing

	if err := database.DB.First(&borrowing, id).Error; err != nil {
		return models.Borrowing{}, err
	}
	return borrowing, nil
}

func (ur *BorrowingRepositoryImp) Update(returningRequest models.ReturningRequest, id string) (models.Borrowing, error) {
	borrowing, err := ur.GetById(id)

	if err != nil {
		return models.Borrowing{}, err
	}

	borrowing.ReturnedAt = returningRequest.ReturnedAt

	var borrowedBook models.Book

	if err := database.DB.First(&borrowedBook, "id = ?", borrowing.BookID).Error; err != nil {
		return models.Borrowing{}, err
	}

	borrowedBook.Quantity += 1

	if err := database.DB.Save(&borrowing).Error; err != nil {
		return models.Borrowing{}, err
	}

	if err := database.DB.Save(&borrowedBook).Error; err != nil {
		return models.Borrowing{}, err
	}

	return borrowing, nil
}

func (ur *BorrowingRepositoryImp) Delete(id string) error {
	borrowing, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&borrowing).Error; err != nil {
		return err
	}

	return nil
}
