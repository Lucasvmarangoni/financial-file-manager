package services

import (
	"encoding/json"
	

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	errors "github.com/Lucasvmarangoni/logella/err"
)

func (u *UserService) Create(name, lastName, cpf, email, password string) error {

	newUser, err := entities.NewUser(name, lastName, cpf, email, password)
	if err != nil {
		return errors.ErrCtx(err, "entities.NewUser")
	}

	err = u.CheckIfUserAlreadyExists(newUser.HashEmail, newUser.HashCPF, nil)
	if err != nil {
		return errors.ErrCtx(err, "u.CheckIfUserAlreadyExists")
	}

	err = u.encrypt(newUser)
	if err != nil {
		return errors.ErrCtx(err, "u.encrypt")
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		return errors.ErrCtx(err, "json.Marshal")
	}

	err = u.RabbitMQ.Publish(string(userJSON), "application/json", config.GetEnvString("rabbitMQ", "exchange"), config.GetEnvString("rabbitMQ", "queue_user"), config.GetEnvString("rabbitMQ", "routingkey_userCreate"))
	if err != nil {
		return errors.ErrCtx(err, "RabbitMQ.Publish")
	}

	u.setToMemcacheIfNotNil(newUser)
	u.Memcached_1.SetUnique(newUser.HashCPF)
	u.Memcached_1.SetUnique(newUser.HashEmail)

	return nil
}


