package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (u *UserService) Update(id, name, lastName, email, password string) error {

	user, err := u.FindById(id, nil)
	if err != nil {
		return errors.NewError(err, "u.FindById")
	}

	if name == "" {
		name = user.Name
	}
	if lastName == "" {
		lastName = user.LastName
	}
	if email == "" {
		email = user.Email
	}
	if password == "" {
		password = user.Password
	}

	newUser, err := entities.NewUser(name, lastName, user.CPF, email, password, user.Admin)
	if err != nil {
		return errors.NewError(err, "entities.NewUser")
	}
	newUser.Update(user.ID, user.CreatedAt)

	newUser, err = u.Repository.Update(newUser, context.Background())
	if err != nil {
		return errors.NewError(err, "Repository.Update")
	}
	log.Info().Str("context", "UserHandler").Msgf("User updated successfully (%s)", user.ID)
	return nil
}
