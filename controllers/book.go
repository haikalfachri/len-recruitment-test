package controllers

import (
	"len-test/models"
	"len-test/services"
	"net/http"

	"github.com/labstack/echo"
)

type BookController struct {
	service services.BookService
}

func InitBookContoller() *BookController {
	return &BookController{
		service: services.InitBookService(),
	}
}

func (uc *BookController) Create(c echo.Context) error {
	var bookRequest models.BookRequest

	c.Bind(&bookRequest) 

	err := bookRequest.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "models invalid",
			Error	:  err.Error(),
		})
	}

	book, err := uc.service.Create(bookRequest)
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
		Data	:  book,
	})
}

func (uc *BookController) GetAll(c echo.Context) error {
	books, err := uc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch all book",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to fetch all book",
		Data	:  books,
	})
}

func (uc *BookController) GetById(c echo.Context) error {
	id := c.Param("id")

	book, err := uc.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to fetch a book by id",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to fetch a book by id",
		Data	:  book,
	})
}

func (uc *BookController) Update(c echo.Context) error {
	var bookRequest models.BookRequest

	c.Bind(&bookRequest)

	err := bookRequest.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "models invalid",
			Error	:  err.Error(),
		})
	}

	id := c.Param("id")
	book, err := uc.service.Update(bookRequest, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to update a book",
			Error	:  err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to update a book",
		Data	:  book,
	})
}

func (uc *BookController) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Status 	: "failed",
			Message	: "failed to delete a book",
			Error	:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[any]{
		Status 	: "success",
		Message	: "success to delete a book",
	})
}

