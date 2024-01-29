package services

import (
	"context"
	"encoding/json"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/logella/err"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func (u *UserService) Create(name, lastName, cpf, email, password string) error {
	newUser, err := entities.NewUser(name, lastName, cpf, email, password, false)
	if err != nil {
		return errors.ErrCtx(err, "entities.NewUser")
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		return errors.ErrCtx(err, "json.Marshal")
	}

	u.RabbitMQ.Publish(string(userJSON), "application/json", config.GetEnv("rabbitMQ_exchange").(string), config.GetEnv("rabbitMQ_routingkey_userCreate").(string))	

	returnChannel := make(chan error)
	go func() {
		err := u.CreateManagement(u.MessageChannel)
		if err != nil {
			returnChannel <- err					
		}		
		returnChannel <- nil	
	}()	
	err = <-returnChannel	
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) CreateManagement(messageChannel chan amqp.Delivery) error {

	u.RabbitMQ.Consume(messageChannel, config.GetEnv("rabbitMQ_routingkey_userCreate").(string))

	for message := range messageChannel {
		var user entities.User
		err := json.Unmarshal(message.Body, &user)
		if err != nil {
			return errors.ErrCtx(err, "json.Unmarshal")
		}

		err = u.Repository.Insert(&user, context.Background())
		if err != nil {
			return errors.ErrCtx(err, "Repository.Insert")
		}
		log.Info().Str("context", "UserHandler").Msgf("User created successfully (%s)", user.ID)
	}
	return nil
}
