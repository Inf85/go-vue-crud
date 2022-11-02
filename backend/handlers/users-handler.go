package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go-vue-grud/backend/helpers"
	repository "go-vue-grud/backend/repository/users"
	"go-vue-grud/backend/services/users"
	"io/ioutil"
	"net/http"
)

var logger *logrus.Logger

func init() {
	logger = helpers.GetLogger()
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userRepository := repository.NewUserRepository()
	userService := users.NewUserService(userRepository)

	allUsers := userService.GetUserList()
	logger.Println(allUsers)
	json.NewEncoder(w).Encode(allUsers)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userRepository := repository.NewUserRepository()
	userService := users.NewUserService(userRepository)

	params := mux.Vars(r)
	userId := params["id"]
	body, _ := ioutil.ReadAll(r.Body)
	res := userService.UpdateUserByID(body, userId)
	if res != nil {
		var msg map[string]string = map[string]string{"Result": "Error", "Message": "User with ID = " + userId + " was not updated"}
		resultJSON, _ := json.Marshal(msg)
		fmt.Fprint(w, string(resultJSON))
	}
	var msg map[string]string = map[string]string{"Result": "Success", "Message": "User with ID = " + userId + " was updated"}
	resultJSON, _ := json.Marshal(msg)
	fmt.Fprint(w, string(resultJSON))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["id"]
	userRepository := repository.NewUserRepository()
	userService := users.NewUserService(userRepository)
	res := userService.DeleteUserById(userId)

	if res != nil {
		var msg map[string]string = map[string]string{"Result": "Error", "Message": "User with ID = " + userId + " was not deleted"}
		resultJSON, _ := json.Marshal(msg)
		fmt.Fprint(w, string(resultJSON))
	}
	var msg map[string]string = map[string]string{"Result": "Success", "Message": "User with ID = " + userId + " was deleted"}
	resultJSON, _ := json.Marshal(msg)
	fmt.Fprint(w, string(resultJSON))
}
