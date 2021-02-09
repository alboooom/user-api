package models

type User struct {
	Id int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}

func (p *Users) appendUs(us *User) {
	p.Users = append(p.Users, *us)
}
