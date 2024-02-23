package entities_test

import (
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/stretchr/testify/assert"
)

var name = "John"
var lastName = "Doe"
var cpf = "990.000.888-00"
var email = "john@gmail.com"
var password = "1234asd$AS"

func TestNewUser(t *testing.T) {

	t.Run("should return a user when provided valid values", func(t *testing.T) {

		user, err := entities.NewUser(name, lastName, cpf, email, password)
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.ID)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, lastName, user.LastName)
		assert.Equal(t, cpf, user.CPF)
	})

	t.Run("should not return a user when provided invalid name", func(t *testing.T) {
		name := "Jo"

		user, err := entities.NewUser(name, lastName, cpf, email, password)
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "name: Jo does not validate as length(3|10)")

	})

	t.Run("should not return a user when provided invalid lastName", func(t *testing.T) {

		lastName := "Do"

		user, err := entities.NewUser(name, lastName, cpf, email, password)
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "last_name: Do does not validate as length(3|50)")
	})

	t.Run("should not return a user when provided invalid cpf", func(t *testing.T) {

		cpf := "990.001.8sss8802"

		user, err := entities.NewUser(name, lastName, cpf, email, password)
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "cpf: 990.001.8sss8802 does not validate as matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)")
	})

	t.Run("should not return a user when provided invalid password", func(t *testing.T) {

		password := "invalid"

		user, err := entities.NewUser(name, lastName, cpf, email, password)
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "")
	})
}

func TestUser_ValidatePassword(t *testing.T) {

	t.Run("", func(t *testing.T) {

		user, err := entities.NewUser(name, lastName, cpf, email, password)

		assert.Nil(t, err)
		err = user.ValidateHashPassword(password)
		assert.Nil(t, err)

		err = user.ValidateHashPassword(password + "i")
		assert.NotNil(t, err)

		assert.NotEqual(t, password, user.Password)
	})

}
