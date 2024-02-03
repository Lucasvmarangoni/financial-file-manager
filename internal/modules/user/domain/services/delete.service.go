package services

import (
	"context"

	errors "github.com/Lucasvmarangoni/logella/err"
	logella "github.com/Lucasvmarangoni/logella/err"
)

func (u *UserService) Delete(id string) error {
	err := u.Repository.Delete(id, context.Background())
	if err != nil {
		return logella.ErrCtx(err, "Repository.Update")
	}
	err = u.deleteFromMemcache(id)
	if err != nil {
		return errors.ErrCtx(err, "u.deleteFromMemcache")
	}
	return nil
}
