package service

import (
	"errors"

	"github.com/shreekumar2901/url-shortener/dto"
	"github.com/shreekumar2901/url-shortener/helpers"
	"github.com/shreekumar2901/url-shortener/repository"
)

type UserService struct{}

func (s *UserService) CreateUser(userDTO *dto.UserRequestDTO) (string, error) {

	hashedPassword, err := helpers.HashPassword(userDTO.Password)

	if err != nil {
		return "", err
	}

	userDTO.Password = hashedPassword

	err = repository.CreateUser(userDTO)

	if err != nil {
		return "", err
	}

	return "User Created Successfully", nil

}

func (s *UserService) GetUserByUserName(username string) (dto.UserResponseDTO, error) {

	var userResponseDTO dto.UserResponseDTO
	user, err := repository.GetUserByUserName(username)

	if err != nil {
		return userResponseDTO, err
	}

	userResponseDTO = dto.UserResponseDTO{
		UserName: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	return userResponseDTO, nil
}

func (s *UserService) GetUserIdByUsername(username string) (string, error) {
	id, err := repository.GetUserIdByUsername(username)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *UserService) FindByCredentials(UsernameOrEmail string, Password string) (dto.UserResponseDTO, error) {
	var userResponseDTO dto.UserResponseDTO

	user, err := repository.FindByCredentials(UsernameOrEmail)

	if err != nil {
		return userResponseDTO, err
	}

	if !helpers.VerifyPassword(user.Password, Password) {
		return userResponseDTO, errors.New("unauthorized request")
	}

	userResponseDTO = dto.UserResponseDTO{
		Email:    user.Email,
		UserName: user.Username,
		Role:     user.Role,
	}

	return userResponseDTO, nil
}

func (s *UserService) DeleteUserByUserName(username string) (string, error) {
	err := repository.DeleteUserByUserName(username)

	if err != nil {
		return "", errors.New("some error occurred during delete")
	}

	return "User deleted succesfully", nil
}
