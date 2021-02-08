package internal

import (
	"encoding/json"
	"io/ioutil"
)

type Employee struct {
	Id int
	FirstName, Email string
}

func NewUser(id int, name, email string) {
	data := Employee{
		Id: id,
		FirstName: name,
		Email:    email,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("internal/db/db.json", file, 0644)
}