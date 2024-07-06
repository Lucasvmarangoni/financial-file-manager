package services

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/logella/err"
)

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

	err = u.deleteCache(id, user.HashEmail, user.HashCPF)
	if err != nil {
		return errors.ErrCtx(err, "u.deleteCache")
	}

	u.updateField(&name, user.Name, name, "Name", &newUpdateValues)
	u.updateField(&lastName, user.LastName, lastName, "LastName", &newUpdateValues)
	u.updateField(&email, user.Email, email, "Email", &newUpdateValues)
	u.updateField(&password, user.Password, newPassword, "Password", &newUpdateValues)

	var oldValues []entities.UpdateLog
	oldValues = append(oldValues, user.UpdateLog...)
	oldValues = append(oldValues, newUpdateValues)

	if newPassword != "" {
		password = newPassword
	}

	err = user.Update(oldValues, name, lastName, email, password)
	if err != nil {
		return errors.ErrCtx(err, "user.Update")
	}

	err = u.encrypt(user)
	if err != nil {
		return errors.ErrCtx(err, "u.encrypt")
	}

	err = u.Repository.Update(user, context.Background())
	if err != nil {
		return errors.ErrCtx(err, "Repository.Update")
	}

	u.Memcached_1.SetUnique(user.HashCPF)
	u.Memcached_1.SetUnique(user.HashEmail)
	u.setToMemcacheIfNotNil(user)
	return nil
}

func (u *UserService) updateField(field *string, oldValue string, newValue string, fieldName string, newUpdateValues *entities.UpdateLog) {
	if newValue == "" {
		*field = oldValue
	} else {
		newUpdateValues.OldValues[fieldName] = oldValue
	}
}
