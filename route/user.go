package route

import (
	"livecode/handlers"
	mysql "livecode/pkg"
	"livecode/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/userall", h.GetAllUser)
	e.GET("/user/:id", h.GetUser)
	e.POST("/user", h.CreateUser)
	e.PATCH("/update/:id", h.UpdateUser)
	e.DELETE("/delete/:id", h.DeleteUser)

}
