package services

import "len-test/repositories"
import "len-test/models"

type StudentService struct {
	repository repositories.StudentRepository
}

func InitStudentService() StudentService {
	return StudentService{
		repository: &repositories.StudentRepositoryImp{},
	}
}

func (us *StudentService) Create(studentRequest models.StudentRequest) (models.Student, error) {
	return us.repository.Create(studentRequest)
}

func (us *StudentService) GetAll() ([]models.Student, error) {
	return us.repository.GetAll()
}

func (us *StudentService) GetById(id string) (models.Student, error) {
	return us.repository.GetById(id)
}

func (us *StudentService) Update(studentRequest models.StudentRequest, id string) (models.Student, error) {
	return us.repository.Update(studentRequest, id)
}

func (us *StudentService) Delete(id string) error {
	return us.repository.Delete(id)
}