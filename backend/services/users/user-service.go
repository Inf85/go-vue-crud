package users

import (
	"go-vue-grud/backend/models"
	repository "go-vue-grud/backend/repository/users"
)

type userService struct {
	repo repository.UserQuery
}

func NewUserService(repo repository.UserQuery) *userService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) GetUserList() []models.User {

	users := u.repo.GetUsers()
	return users

}

func (u *userService) UpdateUserByID(body []byte, id string) error {
	err := u.repo.UpdateUser(body, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) DeleteUserById(id string) error {
	err := u.repo.DeleteUser(id)

	if err != nil {
		return err
	}
	return nil
}
