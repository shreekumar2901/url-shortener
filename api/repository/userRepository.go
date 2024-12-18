package repository

import (
	"errors"

	"github.com/shreekumar2901/url-shortener/database"
	"github.com/shreekumar2901/url-shortener/domain"
	"github.com/shreekumar2901/url-shortener/dto"
	"gorm.io/gorm"
)

func CreateUser(userDTO *dto.UserRequestDTO) error {
	user := domain.User{
		Username: userDTO.UserName,
		Password: userDTO.Password,
		Email:    userDTO.Email,
	}

	db := database.Db.DB

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByUserName(username string) (*domain.User, error) {
	db := database.Db.DB
	var user domain.User

	err := db.Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func DeleteUserByUserName(username string) error {
	db := database.Db.DB

	var user domain.User

	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}

	err = db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil

}
