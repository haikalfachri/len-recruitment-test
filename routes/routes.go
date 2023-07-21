package routes

import (
	"len-test/controllers"

	"github.com/labstack/echo"
)

func SetUpRoutes(e *echo.Echo) {
	studentCtrl := controllers.InitStudentContoller()
	bookCtrl := controllers.InitBookContoller()
	borrowingCtrl := controllers.InitBorrowingContoller()

	e.GET("/students", studentCtrl.GetAll)
	e.GET("/students/:id", studentCtrl.GetById)
	e.POST("/students", studentCtrl.Create)
	e.PUT("/students/:id", studentCtrl.Update)
	e.DELETE("/students/:id", studentCtrl.Delete)

	e.GET("/books", bookCtrl.GetAll)
	e.GET("/books/:id", bookCtrl.GetById)
	e.POST("/books", bookCtrl.Create)
	e.PUT("/books/:id", bookCtrl.Update)
	e.DELETE("/books/:id", bookCtrl.Delete)

	e.GET("/borrowings", borrowingCtrl.GetAll)
	e.GET("/borrowings/:id", borrowingCtrl.GetById)
	e.POST("/borrowings", borrowingCtrl.Create)
	e.PUT("/borrowings/:id", borrowingCtrl.Update)
	e.DELETE("/borrowings/:id", borrowingCtrl.Delete)
}