package models

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/alboooom/user-api/internal"

)

func AdUsers (c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	internal.NewUser(u.Id, u.Name, u.Email)
	return c.JSON(http.StatusOK, u)
}