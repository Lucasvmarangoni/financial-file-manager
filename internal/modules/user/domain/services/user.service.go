package services

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"
)

type UserService struct {
	Repository *repositories.UserRepositoryDb
	RabbitMQ   *queue.RabbitMQ
}

func NewUserService(repo *repositories.UserRepositoryDb, rabbitMQ *queue.RabbitMQ) *UserService {
	UserService := &UserService{
		Repository: repo,
		RabbitMQ:   rabbitMQ,
	}
	return UserService
}
