package handler

import (
	"arrow_food_api/models"
	"arrow_food_api/repo"
	"arrow_food_api/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/register/users", RegisterUser).Methods("POST")
	return r
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