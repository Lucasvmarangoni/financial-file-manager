package services

import (
	"encoding/json"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
)

func (u *UserService) Create(name, lastName, cpf, email, password string) error {
	newUser, err := entities.NewUser(name, lastName, cpf, email, password, false)
	if err != nil {
		return errors.NewError(err, "entities.NewUser")
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		return errors.NewError(err, "json.Marshal")
	}

	u.RabbitMQ.Publish(string(userJSON), "application/json", config.GetEnv("rabbitMQ_exchange").(string), config.GetEnv("rabbitMQ_routingkey_userCreate").(string))
	return nil
}
