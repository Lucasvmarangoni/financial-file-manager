package services

import (
	"context"
	errD "errors"

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
	Memcached      cache.Mencacher[*entities.User]
	Memcached_1    cache.Memcached[bool]
	cacheUser      *entities.User
	aesKey         string
	hmacKey        []byte
}

func NewUserService(
	repo repositories.UserRepository,
	rabbitMQ queue.IRabbitMQ,
	messageChannel chan amqp.Delivery,
	memcached cache.Mencacher[*entities.User],
	memcached_1 cache.Memcached[bool],
) *UserService {
	UserService := &UserService{
		Repository:     repo,
		RabbitMQ:       rabbitMQ,
		MessageChannel: messageChannel,
		Memcached:      memcached,
		Memcached_1:    memcached_1,
		aesKey:         config.ReadSecretString(config.GetEnvString("security", "aes_key")),
		hmacKey:        []byte(config.GetEnvString("security", "hmac_key")),
	}
	return UserService
}

func (u *UserService) setToMemcacheIfNotNil(user *entities.User) {
	if user != nil {
		u.Memcached.Set(user.ID.String(), user)
	}
}

func (u *UserService) returnCachedUserIfExists() (*entities.User, bool) {
	if u.cacheUser != nil {
		return u.cacheUser, true
	}
	return nil, false
}

func (u *UserService) deleteUserCache(id string) error {
	if u.cacheUser != nil {
		err := u.Memcached.Delete(id)
		if err != nil {
			return errors.ErrCtx(err, "u.Memcached.Delete")
		}
	}
	return nil
}

func (u *UserService) deleteEmailAndCpfCache(key string) error {
	if u.cacheUser != nil {
		err := u.Memcached_1.Delete(key)
		if err != nil {
			return errors.ErrCtx(err, "u.Memcached_1.Delete")
		}
	}
	return nil
}

func (u *UserService) encrypt(user *entities.User) error {

	var err error
	user.LastName, err = security.Encrypt(user.LastName, u.aesKey)
	if err != nil {
		return errors.ErrCtx(err, "security.Encrypt LastName")
	}
	user.Email, err = security.Encrypt(user.Email, u.aesKey)
	if err != nil {
		return errors.ErrCtx(err, "security.Encrypt Email")
	}
	user.CPF, err = security.Encrypt(user.CPF, u.aesKey)
	if err != nil {
		return errors.ErrCtx(err, "security.Encrypt CPF")
	}
	return nil
}

func (u *UserService) CheckIfUserAlreadyExists(hashEmail, hashCPF string, ctx context.Context) error {

	// lembrar de deletar o cache de verificação junto de quando um usuário for deletado.
	// Verificar se esta armazenando todas as keys ou sobrescrevendo e mantendo apenas a ultima. ** OK **

	e := errD.New("duplicate key value violates unique constraint")

	err := u.Memcached_1.GetUnique(hashEmail)
	if err == nil {
		return errors.ErrCtx(e, "u.Memcached_1.GetUnique(hashEmail)")
	}
	err = u.Memcached_1.GetUnique(hashCPF)
	if err == nil {
		return errors.ErrCtx(e, "u.Memcached_1.GetUnique(hashCPF)")
	}

	check, err := u.Repository.CheckIfUserAlreadyExists(hashEmail, hashCPF, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "u.Repository.CheckIfUserAlreadyExists")
	}
	if check {
		return errors.ErrCtx(e, "u.Repository.CheckIfUserAlreadyExists")
	}
	return nil
}


func (u *UserService) deleteCache(id, HashEmail, HashCPF string) error {

	err := u.deleteUserCache(id)
	if err != nil {
		return errors.ErrCtx(err, "u.deleteFromMemcache (User)")
	}

	err = u.deleteEmailAndCpfCache(HashEmail)
	if err != nil {
		return errors.ErrCtx(err, "u.deleteFromMemcache (Email)")
	}

	err = u.deleteEmailAndCpfCache(HashCPF)
	if err != nil {
		return errors.ErrCtx(err, "u.deleteFromMemcache (CPF)")
	}
	return nil
}
