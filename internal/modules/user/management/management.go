package management

import (
	"context"
	"encoding/json"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"
	logella "github.com/Lucasvmarangoni/logella/err"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type UserManagement struct {
	Repository *repositories.UserRepositoryDb
	RabbitMQ   *queue.RabbitMQ
}

func NewManagement(repository *repositories.UserRepositoryDb, rabbitMQ *queue.RabbitMQ) *UserManagement {
	return &UserManagement{
		Repository: repository,
		RabbitMQ:   rabbitMQ,
	}
}

func (m *UserManagement) CreateManagement(messageChannel chan amqp.Delivery) error {

	m.RabbitMQ.Consume(messageChannel, config.GetEnv("rabbitMQ_routingkey_userCreate").(string))

	for message := range messageChannel {
		var user entities.User
		err := json.Unmarshal(message.Body, &user)
		if err != nil {
			return logella.ErrCtx(err, "json.Unmarshal")
		}

		newUser, err := m.Repository.Insert(&user, context.Background())
		if err != nil {
			return logella.ErrCtx(err, "Repository.Insert")
		}
		log.Info().Str("context", "UserHandler").Msgf("User created successfully (%s)", newUser.ID)
	}
	log.Print("CreateManagement finished")
	return nil
}
