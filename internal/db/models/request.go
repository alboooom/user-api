package models

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func CreateUser(c echo.Context) (err error) {
	u := new(User)
	users := new(Users)
	if err = c.Bind(u); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	users, err = getJson("db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in reading file")
	}
	users.appendUs(u)
	err = saveJson("db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in writing file")
	}
	return c.JSON(http.StatusOK, u)
}

func UpdateUser(c echo.Context) (err error) {
	updateData := new(User)
	if err = c.Bind(updateData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user id")
	}
	users := new(Users)
	users, err = getJson("db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in reading file")
	}
	for i, user := range users.Users {
		if user.Id == intId {
			users.Users[i].Name = updateData.Name
			users.Users[i].Email = updateData.Email
		}

	}
	err = saveJson("db.json", users)
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
	users := new(Users)
	users, err = getJson("db.json", users)
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
	err = saveJson("db.json", users)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error in writing file")
	}
	return c.NoContent(http.StatusOK)
}

func getJson(file string, buf *Users) (*Users, error) {
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteValue, buf)
	return buf, nil
}

func saveJson(file string, buf *Users) error {
	result, _ := json.Marshal(buf)
	err := ioutil.WriteFile(file, result, 0644)
	if err != nil {
		return err
	}
	return nil
}
