package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	errors "github.com/Lucasvmarangoni/logella/err"
	logella "github.com/Lucasvmarangoni/logella/err"
)

func (u *UserService) Delete(id string) error {
	var user *entities.User

	err := u.Repository.Delete(id, context.Background())
	if err != nil {
		return logella.ErrCtx(err, "Repository.Update")
	}

	user, err = u.FindById(id, nil)
	if err != nil {
		return errors.ErrCtx(err, "u.FindById")
	}

	err = u.deleteCache(id, user.HashEmail, user.HashCPF)
	if err != nil {
		return errors.ErrCtx(err, "u.deleteCache")
	}

	return nil
}
