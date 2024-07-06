package management

import (
	"context"
	"encoding/json"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type UserManagement struct {
	Repository *repositories.UserRepositoryDb
}

func NewManagement(repository *repositories.UserRepositoryDb) *UserManagement {
	return &UserManagement{
		Repository: repository,
	}
}

func (m *UserManagement) CreateManagement(messageChannel chan amqp.Delivery) {

	for message := range messageChannel {

		var user entities.User
		err := json.Unmarshal(message.Body, &user)
		if err != nil {
			errors.ErrCtx(err, "json.Unmarshal")
			continue
		}

		err = m.Repository.Insert(&user, context.Background())
		if err != nil {
			errors.ErrCtx(err, "Repository.Insert ")
			continue
		}
		log.Info().Str("context", "UserHandler").Msgf("User created successfully (%s)", user.ID)		
	}
}
