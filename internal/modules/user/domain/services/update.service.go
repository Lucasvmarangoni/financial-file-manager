package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/logella/err"
)

const fieldNamePassword = "Password"

func (u *UserService) Update(id, name, lastName, email, password, newPassword string) error {
	var newUpdateValues entities.UpdateLog
	newUpdateValues.OldValues = make(map[string]interface{})

	user, err := u.FindById(id, nil)
	if err != nil {
		return errors.ErrCtx(err, "u.FindById")
	}

	err = user.ValidateHashPassword(password)
	if err != nil {
		return errors.ErrCtx(err, "user.ValidateHashPassword")
	}

	u.updateField(&name, user.Name, name, "Name", &newUpdateValues)
	u.updateField(&lastName, user.LastName, lastName, "LastName", &newUpdateValues)
	u.updateField(&email, user.Email, email, "Email", &newUpdateValues)
	u.updateField(&password, password, newPassword, fieldNamePassword, &newUpdateValues)

	var oldValues []entities.UpdateLog
	oldValues = append(oldValues, user.UpdateLog...)
	oldValues = append(oldValues, newUpdateValues)

	if newPassword != "" {
		password = newPassword
	}
	newUser, err := entities.NewUser(name, lastName, user.CPF, email, password)
	if err != nil {
		return errors.ErrCtx(err, "entities.NewUser")
	}

	err = u.encrypt(newUser)
	if err != nil {
		return errors.ErrCtx(err, "u.encrypt")
	}

	newUser.Update(oldValues, user.ID, user.CreatedAt)

	err = u.Repository.Update(newUser, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "Repository.Update")
	}
	u.setToMemcacheIfNotNil(newUser)
	return nil
}

func (u *UserService) updateField(field *string, oldValue string, newValue string, fieldName string, newUpdateValues *entities.UpdateLog) {
	if newValue == "" {
		*field = oldValue
	} else {
		if fieldName == fieldNamePassword {
			oldValue = "****"
		}
		newUpdateValues.OldValues[fieldName] = oldValue
	}
}
