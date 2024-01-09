package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
)

func (u *UserService) AuthzAdmin(adminID, id string) error {
	admin, err := u.FindById(adminID, nil)
	if err != nil {
		return errors.NewError(err, "u.FindById")
	}

	if admin.Admin == true {
		err := u.Repository.ToggleAdmin(id, context.Background())
		if err != nil {
			return errors.NewError(err, "u.Repository.ToggleAdmin")
		}
	}
	return nil
}