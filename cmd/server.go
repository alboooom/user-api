package cmd

import (
	"github.com/labstack/echo"
	"github.com/alboooom/user-api/internal/db/models"
)

func StartServer(){
	e := echo.New()
	e.GET("/", models.Hello)
	e.POST("/adduser", models.AdUsers)
	e.Logger.Fatal(e.Start(":1323"))
}