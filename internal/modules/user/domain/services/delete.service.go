package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (u *UserService) Delete(id string) error {
	err := u.Repository.Delete(id, context.Background())
	if err != nil {
		return errors.NewError(err, "Repository.Update")
	}
	log.Info().Str("context", "UserHandler").Msgf("User deleted successfully (%s)", id)
	return nil
}
