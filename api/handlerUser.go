package handler

import (
	"arrow_food_api/models"
	"arrow_food_api/repo"
	"arrow_food_api/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const baseUrl ="/api/users/"

func GetRouterUser() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/{email}", GetUserByEmail).Methods("GET").Name("users")
	r.HandleFunc("/api/register/users", RegisterUser).Methods("POST")
	r.HandleFunc("/api/users/delete/{email}", DeleteUser).Methods("DELETE")
	return r
}

func GetUserByEmail(rw http.ResponseWriter, r *http.Request) {
	response := &models.Response{}
	var path []string
	if strings.HasPrefix(r.URL.Path, baseUrl) {
		path = strings.Split(r.URL.Path[len(baseUrl):], "/")
		fmt.Println(path)
	}
	var result *models.User
	if len(path) > 0 {
	 result,_ =	services.GetUserByEmail(path[0])
	}

	if result == nil {
		response.Contructor(nil,"No se ha encontrado dicho usuario",http.StatusNoContent)
		json.NewEncoder(rw).Encode(response)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	response.Contructor(result,"",http.StatusOK)
	json.NewEncoder(rw).Encode(response)

}

func RegisterUser(rw http.ResponseWriter, r *http.Request) {
	body,err :=  io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw,err.Error(),http.StatusInternalServerError)
		return
	}
	var user *models.User
	json.Unmarshal(body, &user)

	rw.Header().Set("Content-Type", "application/json")
	response := &models.Response{}

	//comprobar que el email no exista
	existUser,_ :=  repo.GetUserByEmail(user.Email)
	fmt.Println(existUser)
	if existUser != nil {
		response.Contructor(nil,"Email ya esta registrado",http.StatusConflict)
		json.NewEncoder(rw).Encode(response)
		return
	}

	// registrar usuario
	errRegister:= services.RegisterUser(user)
	// comporbarq que haya ido todo correcto
	if errRegister != nil {
		http.Error(rw,err.Error(),http.StatusInternalServerError)
		return
	}

	// se envia el correo que se ha registrado el usuario
	template := models.EmailTemplate{Sender: user.Email, Subject:"registro",Body: "se ha registrado correctamente"}
	email :=models.Email{}
	errorEmail:= email.SendMail(&template)

	if errorEmail !=nil {
		http.Error(rw,errorEmail.Error(),http.StatusInternalServerError)
		// deberia eliminar el usuario creado?
		return
	}
	response.Contructor(user,"",http.StatusCreated)
	json.NewEncoder(rw).Encode(response)
}

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