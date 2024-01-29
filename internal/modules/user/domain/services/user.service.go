package services

import (
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"
	"github.com/streadway/amqp"
)

type UserService struct {
	Repository     repositories.UserRepository
	RabbitMQ       queue.IRabbitMQ
	MessageChannel chan amqp.Delivery
	ReturnChannel  chan error
}

func NewUserService(repo repositories.UserRepository, rabbitMQ queue.IRabbitMQ, messageChannel chan amqp.Delivery, returnChannel chan error) *UserService {
	UserService := &UserService{
		Repository:     repo,
		RabbitMQ:       rabbitMQ,
		MessageChannel: messageChannel,
		ReturnChannel:  returnChannel,
	}
	return UserService
}
