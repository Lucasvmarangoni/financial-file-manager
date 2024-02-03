package dto

import (
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"

	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type UserInput struct {
	Name     string `json:"name" valid:"required,alpha,length(3|10)"`
	LastName string `json:"last_name" valid:"required,alpha,length(3|50)"`
	CPF      string `json:"cpf" valid:"required,matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"required"`
}

type AuthenticationInput struct {
	CPF      string `json:"cpf" valid:"required,matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token" valid:"-"`
}

type UserOutput struct {
	ID        pkg_entities.ID      `json:"id" valid:"-"`
	Name      string               `json:"name" valid:"required,alpha,length(3|10)"`
	LastName  string               `json:"last_name" valid:"required,alpha,length(3|50)"`
	CPF       string               `json:"cpf" valid:"required,matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	Email     string               `json:"email" valid:"email"`
	CreatedAt time.Time            `json:"created_at" valid:"-"`
	UpdateLog []entities.UpdateLog `json:"update_log" valid:"-"`
}

type UserUpdateInput struct {
	Name        string `json:"name" valid:"-"`
	LastName    string `json:"last_name" valid:"-"`
	Email       string `json:"email" valid:"-"`
	Password    string `json:"password" valid:"required"`
	NewPassword string `json:"new_password" valid:"-"`
}
