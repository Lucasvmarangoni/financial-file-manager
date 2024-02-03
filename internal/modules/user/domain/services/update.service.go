package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/logella/err"
)

func (u *UserService) Update(id, name, lastName, email, password string) error {
	var oldValues map[string]interface{}
	oldValues = make(map[string]interface{})

	user, err := u.FindById(id, nil)
	if err != nil {
		return errors.ErrCtx(err, "u.FindById")
	}

	u.updateField(&name, user.Name, name, "Name", oldValues)
	u.updateField(&lastName, user.LastName, lastName, "LastName", oldValues)
	u.updateField(&email, user.Email, email, "Email", oldValues)
	u.updateField(&password, user.Password, password, "Password", oldValues)

	newUser, err := entities.NewUser(name, lastName, user.CPF, email, password)
	if err != nil {
		return errors.ErrCtx(err, "entities.NewUser")
	}
	newUser.Update(oldValues, user.ID, user.CreatedAt)

	err = u.Repository.Update(newUser, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "Repository.Update")
	}
	u.setToMemcacheIfNotNil(newUser)
	return nil
}

func (u *UserService) updateField(field *string, oldValue interface{}, newValue string, fieldName string, oldValues map[string]interface{}) {
	if newValue == "" {
		*field = oldValue.(string)
	} else {
		oldValues[fieldName] = oldValue
	}
}
