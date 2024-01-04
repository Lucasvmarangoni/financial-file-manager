package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
)

func (u *UserService) Create(name, lastName, cpf, email, password string, admin bool) (pkg_entities.ID, error) {
	newUser, err := entities.NewUser(name, lastName, cpf, email, password, admin)
	if err != nil {
		return pkg_entities.Nil(), errors.NewError(err, "entities.NewUser")
	}
	newUser, err = u.Repository.Insert(newUser, context.Background())
	if err != nil {
		return pkg_entities.Nil(), errors.NewError(err, "Repository.Insert")
	}
	return newUser.ID, nil
}
