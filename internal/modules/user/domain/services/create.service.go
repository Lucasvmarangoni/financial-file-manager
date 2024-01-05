package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (u *UserService) Create(name, lastName, cpf, email, password string, admin bool) error {
	newUser, err := entities.NewUser(name, lastName, cpf, email, password, admin)
	if err != nil {
		return errors.NewError(err, "entities.NewUser")
	}
	newUser, err = u.Repository.Insert(newUser, context.Background())
	if err != nil {
		return errors.NewError(err, "Repository.Insert")
	}
	log.Info().Str("context", "UserHandler").Msgf("User created successfully (%s)", newUser.ID)
	return nil
}
