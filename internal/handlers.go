package internal


import (
	"encoding/json"
	"fmt"
	"github.com/alboooom/user-api/internal/db/models"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GetUsers(c echo.Context) (err error) {
	u := new(models.Users)
	if err = c.Bind(u); err != nil {
		return err
	}
	jsonFile, err := os.Open("internal/db/db.json")
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
	users := new(models.Users)
	users, err = models.GetJson("internal/db/db.json", users)
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

func CreateUser(c echo.Context) (err error) {
	u := new(models.User)
	users := new(models.Users)
	if err = c.Bind(u); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	users, err = models.GetJson("internal/db/db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in reading file")
	}
	users.AppendUs(u)
	err = models.SaveJson("internal/db/db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in writing file")
	}
	return c.JSON(http.StatusOK, u)
}

func UpdateUser(c echo.Context) (err error) {
	updateData := new(models.User)
	if err = c.Bind(updateData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user id")
	}
	users := new(models.Users)
	users, err = models.GetJson("internal/db/db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in reading file")
	}
	for i, user := range users.Users {
		if user.Id == intId {
			users.Users[i].Name = updateData.Name
			users.Users[i].Email = updateData.Email
		}

	}
	err = models.SaveJson("internal/db/db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in writing file")
	}
	return c.NoContent(http.StatusOK)
}

func DeleteUser(c echo.Context) (err error) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user id")
	}
	users := new(models.Users)
	users, err = models.GetJson("internal/db/db.json", users)
	if err!=nil {
		return c.String(http.StatusBadRequest, "Error in reading file")
	}
	for i, user := range users.Users{
		if user.Id == intId{
			users.Users[i] = users.Users[len(users.Users)-1]
			users.Users = users.Users[:len(users.Users)-1]
			break
		}
	}
	err = models.SaveJson("internal/db/db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in writing file")
	}
	return c.NoContent(http.StatusOK)
}


