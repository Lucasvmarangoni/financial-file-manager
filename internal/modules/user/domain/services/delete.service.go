package services

import (
	"context"

	logella "github.com/Lucasvmarangoni/logella/err"
)

func (u *UserService) Delete(id string) error {
	err := u.Repository.Delete(id, context.Background())
	if err != nil {
		return logella.ErrCtx(err, "Repository.Update")
	}
	return nil
}
