package models

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GetUsers(c echo.Context) (err error) {
	u := new(Users)
	if err = c.Bind(u); err != nil {
		return err
	}
	jsonFile, err := os.Open("db.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(byteValue, u)
	var result string
	for _, user:=range u.Users{
		id := strconv.Itoa(user.Id)
		result += id +  "| name: " + user.Name + ",  email: " + user.Email  + "\n"
	}
	return c.String(http.StatusOK, result)
}

func GetUser(c echo.Context) error{
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user id")
	}
	users := new(Users)
	users, err = getJson("db.json", users)
	if err != nil{
		return c.String(http.StatusBadRequest, "Error in reading file")
	}
	for _, user := range users.Users{
		if user.Id == intID{
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

