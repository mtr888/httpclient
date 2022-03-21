package gorest

import "fmt"

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

func NewUser(name string, email string, gender string, status string) *User {
	return &User{
		Name:   name,
		Email:  email,
		Gender: gender,
		Status: status,
	}
}

func (u *User) Info() string {
	return fmt.Sprintf("[Id] %d | [Name] %s | [Email] %s | [Gender] %s | [Status] %s", u.Id, u.Name, u.Email, u.Gender, u.Status)
}

type usersResponse []User
