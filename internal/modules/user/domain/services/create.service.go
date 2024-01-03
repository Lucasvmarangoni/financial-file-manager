package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	logger "github.com/Lucasvmarangoni/financial-file-manager/pkg/log"
)

type CreateService struct {
	Repository *repositories.UserRepositoryDb
}

func NewCreateService(repo *repositories.UserRepositoryDb) *CreateService {
	createService := &CreateService{
		Repository: repo,
	}
	return createService
}
func (c *CreateService) Create(name, lastName, cpf, email, password string, admin bool) error {
	newUser, err := entities.NewUser(name, lastName, cpf, email, password, admin)
	if err != nil {
		return logger.NewError(err, "entities.NewUser")
	}
	newUser, err = c.Repository.Insert(newUser, context.Background())
	if err != nil {
		return logger.NewError(err, "Repository.Insert")
	}
	return nil
}
