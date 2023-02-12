package handler

import (
	"arrow_food_api/models"
	"arrow_food_api/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const baseUrl ="/api/users/"


func DeleteUser(rw http.ResponseWriter, r *http.Request){
	response := &models.Response{}
	var path []string
	if strings.HasPrefix(r.URL.Path, baseUrl) {
		path = strings.Split(r.URL.Path[len(baseUrl):], "/")
		fmt.Println(path)
	}
	var user *models.User
	if len(path) > 0 {
		user,_ =	services.GetUserByEmail(path[1])
	}
	if user == nil {
		response.Contructor(nil,"El usuario no exist", http.StatusNotFound)
		json.NewEncoder(rw).Encode(response)
		return
	}
	err := services.DeleteUserByEmail(user)

	if err != nil {
		response.Contructor(nil,"No se pudo borrar dicho usuario", http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(response)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	response.Contructor(user,"",http.StatusNoContent)
	json.NewEncoder(rw).Encode(response)
}