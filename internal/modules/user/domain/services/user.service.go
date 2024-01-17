package services

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"
)

type UserService struct {
	Repository repositories.UserRepository
	RabbitMQ   queue.IRabbitMQ
}

func NewUserService(repo repositories.UserRepository, rabbitMQ queue.IRabbitMQ) *UserService {
	UserService := &UserService{
		Repository: repo,
		RabbitMQ:   rabbitMQ,
	}
	return UserService
}
