package internal

import (
	"fmt"
	"github.com/alboooom/user-api/internal/db/models"
	"os"
)


func GetAllUsers(u  *models.User ) {
	jsonFile, err := os.Open("/internal/db/db.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	//file, _ := json.MarshalIndent(jsonFile, "", " ")

	//_ = ioutil.WriteFile("internal/db/db.json", file, 0644)
}