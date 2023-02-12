package handler

import (
	"arrow_food_api/models"
	"arrow_food_api/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)



func GetUserByEmail(rw http.ResponseWriter, r *http.Request) {
	response := &models.Response{}
	var path []string
	if strings.HasPrefix(r.URL.Path, baseUrl) {
		path = strings.Split(r.URL.Path[len(baseUrl):], "/")
		fmt.Println(path)
	}
	var result *models.User
	if len(path) > 0 {
		result, _ = services.GetUserByEmail(path[0])
	}

	if result == nil {
		response.Contructor(nil, "No se ha encontrado dicho usuario", http.StatusNoContent)
		json.NewEncoder(rw).Encode(response)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	response.Contructor(result, "", http.StatusOK)
	json.NewEncoder(rw).Encode(response)

}