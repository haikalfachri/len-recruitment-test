package services

import "len-test/repositories"
import "len-test/models"

type BorrowingService struct {
	repository repositories.BorrowingRepository
}

func InitBorrowingService() BorrowingService {
	return BorrowingService{
		repository: &repositories.BorrowingRepositoryImp{},
	}
}

func (us *BorrowingService) Create(borrowingRequest models.BorrowingRequest) (models.Borrowing, error) {
	return us.repository.Create(borrowingRequest)
}

func (us *BorrowingService) GetAll() ([]models.Borrowing, error) {
	return us.repository.GetAll()
}

func (us *BorrowingService) GetById(id string) (models.Borrowing, error) {
	return us.repository.GetById(id)
}

func (us *BorrowingService) Update(returningRequest models.ReturningRequest, id string) (models.Borrowing, error) {
	return us.repository.Update(returningRequest, id)
}

func (us *BorrowingService) Delete(id string) error {
	return us.repository.Delete(id)
}

