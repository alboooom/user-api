package models

type User struct {
	Id int `json:"id" form:"id" query:"id"`
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`

}
