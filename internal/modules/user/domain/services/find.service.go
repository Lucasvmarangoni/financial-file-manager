package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
)


func (u *UserService) FindByEmail(email string, ctx context.Context) (*entities.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.Repository.FindByEmail(email, ctx)
	if err != nil {
		return nil, errors.NewError(err, "Repository.FindByEmail")
	}
	return user, nil
}

func (u *UserService) FindById(id string, ctx context.Context) (*entities.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	parsedId, err := pkg_entities.ParseID(id)
	if err != nil {
		return nil, errors.NewError(err, "pkg_entities.ParseID")
	}

	user, err := u.Repository.FindById(parsedId, ctx)
	if err != nil {
		return nil, errors.NewError(err, "Repository.FindById")
	}
	return user, nil
}

func (u *UserService) FindByCpf(cpf string, ctx context.Context) (*entities.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.Repository.FindByCpf(cpf, ctx)
	if err != nil {
		return nil, errors.NewError(err, "Repository.FindByCpf")
	}
	return user, nil
}
