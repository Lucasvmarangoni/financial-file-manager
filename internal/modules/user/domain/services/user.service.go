package services

import (
	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/cache"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/queue"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/security"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/streadway/amqp"
)

type UserService struct {
	Repository     repositories.UserRepository
	RabbitMQ       queue.IRabbitMQ
	MessageChannel chan amqp.Delivery
	ReturnChannel  chan error
	memcached      cache.Mencacher[*entities.User]
	cacheUser      *entities.User
}

func NewUserService(
	repo repositories.UserRepository,
	rabbitMQ queue.IRabbitMQ,
	messageChannel chan amqp.Delivery,
	returnChannel chan error,
	memcached cache.Mencacher[*entities.User],
) *UserService {
	UserService := &UserService{
		Repository:     repo,
		RabbitMQ:       rabbitMQ,
		MessageChannel: messageChannel,
		ReturnChannel:  returnChannel,
		memcached:      memcached,
	}
	return UserService
}

func (u *UserService) setToMemcacheIfNotNil(user *entities.User) {
	if user != nil {
		u.memcached.Set(user.ID.String(), user)
	}
}

func (u *UserService) returnCachedUserIfExists() (*entities.User, bool) {
	if u.cacheUser != nil {
		return u.cacheUser, true
	}
	return nil, false
}

func (u *UserService) deleteFromMemcache(id string) error {
	if u.cacheUser != nil {
		err := u.memcached.Delete(id)
		if err != nil {
			return errors.ErrCtx(err, "u.memcached.Delete")
		}
	}
	return nil
}

func (u *UserService) encrypt(user *entities.User) error {
	aes_key := config.GetEnvString("security", "aes_key")
	var err error

	user.LastName, err = security.Encrypt(user.LastName, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Encrypt LastName")
	}
	user.Email, err = security.Encrypt(user.Email, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Encrypt Email")
	}
	user.CPF, err = security.Encrypt(user.CPF, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Encrypt CPF")
	}
	return nil
}
