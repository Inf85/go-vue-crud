package repository

import (
	"encoding/json"
	"go-vue-grud/backend/database"
	"go-vue-grud/backend/models"
)

type userQuery struct {
}

type UserQuery interface {
	GetUsers() []models.User
	UpdateUser([]byte, string) error
	DeleteUser(string) error
}

func NewUserRepository() *userQuery {
	return &userQuery{}
}

func (u *userQuery) GetUsers() []models.User {
	var users []models.User
	result, err := database.DB.Query("SELECT id, first_name," +
		"last_name,email from users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var user models.User
		err := result.Scan(&user.ID, &user.FirstName,
			&user.LastName, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	return users
}

func (u *userQuery) UpdateUser(body []byte, id string) error {
	stmt, err := database.DB.Prepare("UPDATE users SET first_name = ?," +
		"last_name= ?, email=? WHERE id = ?")
	if err != nil {
		return err
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	first_name := keyVal["firstName"]
	last_name := keyVal["lastName"]
	email := keyVal["email"]

	_, err = stmt.Exec(first_name, last_name, email,
		id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userQuery) DeleteUser(id string) error {
	stmt, err := database.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
