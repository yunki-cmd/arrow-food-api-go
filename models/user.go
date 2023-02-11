package models

import "fmt"

type User struct {
	id string
	Nombre   string `json:"nombre"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateUser(nombre string, password string, email string) (*User, error) {
	user := User{Nombre: nombre, Password: password, Email: email}
	return &user, nil
}

func (user *User) ToString() string {
	return fmt.Sprintf("user: %v %v %v \n", user.Nombre, user.Password, user.Email)
}

func (user *User) SetNombre(nombre string){
	user.Nombre = nombre
}

func (user *User) SetPassword(password string){
	user.Password = password
}

func (user *User) SetEmail(email string){
	user.Email = email
}

func (user *User) GetNombre() string{
	return user.Nombre
}

func (user *User) GetPassword() string{
	return user.Password
}

func (user *User) GetEmail() string{
	return user.Email
}