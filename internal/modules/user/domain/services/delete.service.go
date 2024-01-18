package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
)

func (u *UserService) Delete(id string) error {
	err := u.Repository.Delete(id, context.Background())
	if err != nil {
		return errors.NewError(err, "Repository.Update")
	}
	return nil
}
