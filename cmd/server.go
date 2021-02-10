package cmd

import (
	"github.com/alboooom/user-api/internal"
	"github.com/labstack/echo"
)

func StartServer(){
	e := echo.New()
	e.GET("/users", internal.GetUsers)
	e.POST("/users/new", internal.CreateUser)
	e.GET("/users/:id", internal.GetUser)
	e.PUT("/users/:id", internal.UpdateUser)
	e.DELETE("/users/:id", internal.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}