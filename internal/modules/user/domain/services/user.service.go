package services

import "github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"

type UserService struct {
	Repository *repositories.UserRepositoryDb
}

func NewUserService(repo *repositories.UserRepositoryDb) *UserService {
	UserService := &UserService{
		Repository: repo,
	}
	return UserService
}