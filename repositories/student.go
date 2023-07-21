package repositories

import (
	"len-test/database"
	"len-test/models"
)

type StudentRepositoryImp struct {
}

func InitStudentRepository() StudentRepository {
	return &StudentRepositoryImp{}
}

func (ur *StudentRepositoryImp) Create(studentRequest models.StudentRequest) (models.Student, error) {
	var student models.Student = models.Student{
		Name: studentRequest.Name,
		SID: studentRequest.SID,
		Major: studentRequest.Major,
	}

	if err := database.DB.Create(&student).Error; err != nil {
		return models.Student{}, err
	}

	if err := database.DB.Last(&student).Error; err != nil {
		return models.Student{}, err
	}

    return student, nil
}

func (ur *StudentRepositoryImp) GetAll() ([]models.Student, error) {
	var students []models.Student

	if err := database.DB.Find(&students).Error; err != nil {
		return students, err
	}
	return students, nil
}

func (ur *StudentRepositoryImp) GetById(id string) (models.Student, error) {
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func (ur *StudentRepositoryImp) Update(studentRequest models.StudentRequest, id string) (models.Student, error) {
	student, err := ur.GetById(id)

	if err != nil {
		return models.Student{}, err
	}

	if err := database.DB.Save(&student).Error; err != nil {
		return models.Student{}, err
	}

    return student, nil
}

func (ur *StudentRepositoryImp) Delete(id string) error {
	student, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&student).Error; err != nil {
		return err
	}

    return nil
}