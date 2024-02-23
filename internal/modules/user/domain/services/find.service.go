package services

import (
	"context"
	"encoding/json"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/security"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/rs/zerolog/log"
)

func (u *UserService) FindById(id string, ctx context.Context) (*entities.User, error) {

	u.cacheUser = u.getCache(id)
	if cachedUser, ok := u.returnCachedUserIfExists(); ok {
		err := u.decrypt(cachedUser)
		if err != nil {
			return nil, errors.ErrCtx(err, "u.decrypt")
		}
		return cachedUser, nil
	}

	if ctx == nil {
		ctx = context.Background()
	}

	parsedId, err := pkg_entities.ParseID(id)
	if err != nil {
		return nil, errors.ErrCtx(err, "pkg_entities.ParseID")
	}
	user, err := u.Repository.FindById(parsedId, ctx)
	if err != nil {
		return nil, errors.ErrCtx(err, "Repository.FindById")
	}

	err = u.decrypt(user)
	if err != nil {
		return nil, errors.ErrCtx(err, "u.decrypt")
	}

	u.setToMemcacheIfNotNil(user)
	return user, nil
}

func (u *UserService) FindByEmail(hashEmail string, ctx context.Context) (*entities.User, error) {

	if cachedUser, ok := u.returnCachedUserIfExists(); ok {
		return cachedUser, nil
	}

	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.Repository.FindByEmail(hashEmail, ctx)
	if err != nil {
		return nil, errors.ErrCtx(err, "Repository.FindByEmail")
	}
	u.setToMemcacheIfNotNil(user)
	return user, nil
}

func (u *UserService) FindByCpf(hashCPF string, ctx context.Context) (*entities.User, error) {

	if cachedUser, ok := u.returnCachedUserIfExists(); ok {
		return cachedUser, nil
	}

	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.Repository.FindByCpf(hashCPF, ctx)
	if err != nil {
		return nil, errors.ErrCtx(err, "Repository.FindByCpf")
	}
	u.setToMemcacheIfNotNil(user)
	return user, nil
}

func (u *UserService) getCache(key string) *entities.User {
	item, err := u.memcached.Get(key)
	if err == nil && item != nil {
		var cachedUser entities.User
		err = json.Unmarshal(item.Value, &cachedUser)
		if err == nil {
			return &cachedUser
		} else {
			log.Error().Err(errors.ErrCtx(err, "json.Unmarshal")).Msg("Get User cache error")
		}
	}
	return nil
}

func (u *UserService) decrypt(user *entities.User) error {
	aes_key := config.GetEnv("security_aes_key").(string)
	var err error

	user.LastName, err = security.Decrypt(user.LastName, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Decrypt LastName")
	}
	user.Email, err = security.Decrypt(user.Email, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Decrypt Email")
	}
	user.CPF, err = security.Decrypt(user.CPF, aes_key)
	if err != nil {
		return errors.ErrCtx(err, "security.Decrypt CPF")
	}	
	return nil
}
