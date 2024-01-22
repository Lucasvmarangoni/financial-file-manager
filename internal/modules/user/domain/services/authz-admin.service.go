package services

import (
	"context"

	logella "github.com/Lucasvmarangoni/logella/err"
)

func (u *UserService) AuthzAdmin(adminID, id string) error {
	admin, err := u.FindById(adminID, nil)
	if err != nil {
		return logella.ErrCtx(err, "u.FindById")
	}

	if admin.Admin == true {
		err := u.Repository.ToggleAdmin(id, context.Background())
		if err != nil {
			return logella.ErrCtx(err, "u.Repository.ToggleAdmin")
		}
	}
	return nil
}
