package services

import (
	"arrow_food_api/models"
	"arrow_food_api/repo"
	"fmt"
)

func GetUserByEmail(email string) (*models.User, error){

	user,err := repo.GetUserByEmail(email)

	if err != nil {
		fmt.Println("error getting user by email: ",email)
		return nil,err
	}
	return user,nil
}

func RegisterUser(user *models.User) error {
	err := repo.RegisterUser(user)
	if err != nil {
		fmt.Println("error insert user : ", user)
		return err
	}
	return nil
}


func DeleteUserByEmail(user *models.User) error {
	err := repo.DeleteUser(user)
	if err != nil {
		fmt.Println("error deleting user : ", user)
		return err
	}
	return nil
}

func LoginUserByUsername(username string, password string) (*models.User,bool, error){

	user,err :=repo.GetUserByUsername(username)

	if err != nil {
		return nil,false,err
	}

	if user.Nombre == username && user.Password == password {
		return user,true, nil
	}

	return nil,false,err

}