package services

import (
	"context"

	"github.com/Lucasvmarangoni/logella/err"
)

func (u *UserService) AuthzAdmin(id string) error {

	err := u.Repository.ToggleAdmin(id, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "u.Repository.ToggleAdmin")
	}

	return nil
}
