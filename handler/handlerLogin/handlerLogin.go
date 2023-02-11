package handlerlogin

import (
	"arrow_food_api/models"
	"arrow_food_api/services"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)


type UserLogin struct {
	Nombre string `json:"nombre"`
	Password string `json:"password"`
}

func Login(rw http.ResponseWriter, r *http.Request) {
	body,_ := io.ReadAll(r.Body)
	
	decode :=json.NewDecoder(strings.NewReader(string(body)))
	
	var response  models.Response
	var data map[string]interface{}

	err :=decode.Decode(&data)

	if err != nil {
		panic("error decoding")
	}

	rw.Header().Set("content-type", "application/json")

	user,isLogin,_ :=services.LoginUserByUsername(data["nombre"].(string),data["password"].(string))
	resutl := make(map[string]interface{})

	if isLogin && user != nil {
		resutl["email"] = user.Email
		resutl["user"] = isLogin
		response.Contructor(resutl,"",http.StatusOK)
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(resutl)
		return
	}
	resutl["email"]=""
	resutl["user"] =isLogin
	response.Contructor(resutl,"error login",http.StatusUnauthorized)
	rw.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(rw).Encode(resutl)
	//s,_ := json.Marshal(data)


	/* var data UserLogin
	json.Unmarshal(body, &data)
	fmt.Println(data.Nombre)
	fmt.Println(data.Password) */
	
	

}
