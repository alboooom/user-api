package cmd

import (
	"github.com/alboooom/user-api/internal/db/models"
	"github.com/labstack/echo"
)

func StartServer(){
	e := echo.New()
	e.GET("/users", models.GetUsers)
	e.POST("/users/new", models.CreateUser)
	e.GET("/users/:id", models.GetUser)
	e.PUT("/users/:id", models.UpdateUser)
	e.DELETE("/users/:id", models.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}