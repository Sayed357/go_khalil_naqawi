package services

import (
	"client-data/models"
	"client-data/repositories"
	"client-data/utils"
)

type UserService struct {
	repository repositories.UserRepository
	jwtOptions models.JWTOptions
}

func InitUserService(jwtOptions models.JWTOptions) UserService {
	return UserService{
		repository: &repositories.UserRepositoryImpl{},
		jwtOptions: jwtOptions,
	}
}

func (us *UserService) Register(registerReq models.RegisterRequest) (models.User, error) {
	return us.repository.Register(registerReq)
}

func (us *UserService) Login(loginReq models.LoginRequest) (string, error) {
	user, err := us.repository.GetByEmail(loginReq)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(int(user.ID), us.jwtOptions)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) GetUserInfo(id string) (models.User, error) {
	return us.repository.GetUserInfo(id)
}
