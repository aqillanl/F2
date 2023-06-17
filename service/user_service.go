package service

import (
	"finalproject2/model"
	"finalproject2/repository"
)

type UserService interface {
	RegisterUser(user model.User) (model.User, error)
	LoginUser(user model.User) (model.User, error)
	UpdateUser(id int, user model.UserUpdateRequest) (model.User, error)
	DeleteUser(id int) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) RegisterUser(newUser model.User) (model.User, error) {

	user := model.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
		Age:      newUser.Age,
	}

	newUser, err := s.repository.RegisterUser(user)

	return newUser, err
}

func (s *userService) LoginUser(user model.User) (model.User, error) {
	existingUser, err := s.repository.LoginUser(user)
	return existingUser, err
}

func (s *userService) UpdateUser(id int, userUpdate model.UserUpdateRequest) (model.User, error) {
	user, _ := s.repository.ShowUser(id)

	user.Email = userUpdate.Email
	user.Username = userUpdate.Username

	updateUser, err := s.repository.UpdateUser(user)
	return updateUser, err
}

func (s *userService) DeleteUser(id int) error {

	user, _ := s.repository.ShowUser(id)

	err := s.repository.DeleteUser(user)
	return err
}
