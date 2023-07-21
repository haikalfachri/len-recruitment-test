package controllers

import (
	"len-test/models"
	"len-test/services"
	"net/http"

	"github.com/labstack/echo"
)

type StudentController struct {
	service services.StudentService
}

func InitStudentContoller() *StudentController {
	return &StudentController{
		service: services.InitStudentService(),
	}
}

func (uc *StudentController) Create(c echo.Context) error {
	var studentRequest models.StudentRequest

	c.Bind(&studentRequest) 

	err := studentRequest.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "models invalid",
			Error	:  err.Error(),
		})
	}

	student, err := uc.service.Create(studentRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "request invalid",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "create success",
		Data	:  student,
	})
}

func (uc *StudentController) GetAll(c echo.Context) error {
	students, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all student",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all student",
		Data	:  students,
	})
}

func (uc *StudentController) GetById(c echo.Context) error {
	id := c.Param("id")

	student, err := uc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a student by id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a student by id",
		Data	:  student,
	})
}

func (uc *StudentController) Update(c echo.Context) error {
	var studentRequest models.StudentRequest

	c.Bind(&studentRequest)

	err := studentRequest.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "models invalid",
			Error	:  err.Error(),
		})
	}

	id := c.Param("id")
	student, err := uc.service.Update(studentRequest, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to update a student",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to update a student",
		Data	:  student,
	})
}

func (uc *StudentController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to delete a student",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to delete a student",
	})
}

