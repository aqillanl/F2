package repository

import (
	"finalproject2/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	ShowUser(id int) (model.User, error)
	RegisterUser(user model.User) (model.User, error)
	LoginUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) ShowUser(id int) (model.User, error) {
	var user model.User
	err := r.db.Preload("User").Find(&user, id).Error
	return user, err
}

func (r *userRepository) RegisterUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) LoginUser(user model.User) (model.User, error) {
	err := r.db.Debug().Where("email = ?", user.Email).Take(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(user model.User) (model.User, error) {
	err := r.db.Debug().Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteUser(user model.User) error {
	err := r.db.Debug().Delete(&user).Error
	if err == nil {
		err = r.db.Debug().Delete(&user).Error
	}
	return err
}
