package controllers

import (
	"len-test/models"
	"len-test/services"
	"net/http"

	"github.com/labstack/echo"
)

type BorrowingController struct {
	service services.BorrowingService
}

func InitBorrowingContoller() *BorrowingController {
	return &BorrowingController{
		service: services.InitBorrowingService(),
	}
}

func (uc *BorrowingController) Create(c echo.Context) error {
	var borrowingRequest models.BorrowingRequest

	c.Bind(&borrowingRequest) 

	err := borrowingRequest.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "models invalid",
			Error	:  err.Error(),
		})
	}

	borrowing, err := uc.service.Create(borrowingRequest)
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
		Data	:  borrowing,
	})
}

func (uc *BorrowingController) GetAll(c echo.Context) error {
	borrowings, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all borrowing",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all borrowing",
		Data	:  borrowings,
	})
}

func (uc *BorrowingController) GetById(c echo.Context) error {
	id := c.Param("id")

	borrowing, err := uc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a borrowing by id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a borrowing by id",
		Data	:  borrowing,
	})
}

func (uc *BorrowingController) Update(c echo.Context) error {
	var returningRequest models.ReturningRequest

	c.Bind(&returningRequest)

	err := returningRequest.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "models invalid",
			Error	:  err.Error(),
		})
	}

	id := c.Param("id")
	borrowing, err := uc.service.Update(returningRequest, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to update a borrowing",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to update a borrowing",
		Data	:  borrowing,
	})
}

func (uc *BorrowingController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to delete a borrowing",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to delete a borrowing",
	})
}

