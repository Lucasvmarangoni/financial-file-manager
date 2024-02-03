package services

import (
	"context"
	"encoding/json"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/rs/zerolog/log"
)

func (u *UserService) FindByEmail(email string, ctx context.Context) (*entities.User, error) {

	if cachedUser, ok := u.returnCachedUserIfExists(); ok {
		return cachedUser, nil
	}

	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.Repository.FindByEmail(email, ctx)
	if err != nil {
		return nil, errors.ErrCtx(err, "Repository.FindByEmail")
	}
	u.setToMemcacheIfNotNil(user)
	return user, nil
}

func (u *UserService) FindById(id string, ctx context.Context) (*entities.User, error) {

	u.cacheUser = u.getCache(id)
	if cachedUser, ok := u.returnCachedUserIfExists(); ok {
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

	u.setToMemcacheIfNotNil(user)
	return user, nil
}

func (u *UserService) FindByCpf(cpf string, ctx context.Context) (*entities.User, error) {

	if cachedUser, ok := u.returnCachedUserIfExists(); ok {
		return cachedUser, nil
	}

	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.Repository.FindByCpf(cpf, ctx)
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
