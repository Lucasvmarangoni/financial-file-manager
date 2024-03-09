package services

import (
	"encoding/json"
	"sync"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	errors "github.com/Lucasvmarangoni/logella/err"
)

func (u *UserService) Create(name, lastName, cpf, email, password string) error {
	newUser, err := entities.NewUser(name, lastName, cpf, email, password)
	if err != nil {
		return errors.ErrCtx(err, "entities.NewUser")
	}

	err = u.encrypt(newUser)
	if err != nil {
		return errors.ErrCtx(err, "u.encrypt")
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		errors.ErrCtx(err, "json.Marshal")
	}

	err = u.RabbitMQ.Publish(string(userJSON), "application/json", config.GetEnvString("rabbitMQ", "exchange"), config.GetEnvString("rabbitMQ", "routingkey_userCreate"))
	if err != nil {
		return errors.ErrCtx(err, "RabbitMQ.Publish")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err = <-u.ReturnChannel
		wg.Done()
	}()
	wg.Wait()

	if err != nil {
		return errors.ErrCtx(err, "CreateManagement")
	}

	return nil
}
