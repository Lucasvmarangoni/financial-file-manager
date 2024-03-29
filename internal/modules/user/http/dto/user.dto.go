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
	Name     string `json:"name" valid:"required,matches(^[a-zA-Z ]+$),length(3|10)"`
	LastName string `json:"last_name" valid:"required,matches(^[a-zA-Z ]+$),length(3|50)"`
	CPF      string `json:"cpf" valid:"matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"required"`
}

type AuthenticationInput struct {
	CPF      string `json:"cpf" valid:"matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"required"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token" valid:"-"`
}

type UserOutput struct {
	ID         pkg_entities.ID      `json:"id" valid:"-"`
	Name       string               `json:"name" valid:"required,matches(^[a-zA-Z ]+$),length(3|10)"`
	LastName   string               `json:"last_name" valid:"required,matches(^[a-zA-Z ]+$),length(3|50)"`
	CPF        string               `json:"cpf" valid:"required,matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	Email      string               `json:"email" valid:"email"`
	OtpEnabled bool                 `json:"otp_enabled" valid:"-"`
	CreatedAt  time.Time            `json:"created_at" valid:"-"`
	UpdateLog  []entities.UpdateLog `json:"update_log" valid:"-"`
}

type UserUpdateInput struct {
	Name        string `json:"name" valid:"-"`
	LastName    string `json:"last_name" valid:"-"`
	Email       string `json:"email" valid:"-"`
	Password    string `json:"password" valid:"required"`
	NewPassword string `json:"new_password" valid:"-"`
}

type OTPInput struct {
	Token string `json:"token" valid:"length(6),numeric"`
}

type OTPOutput struct {
	Base32     string `json:"base32" valid:"required"`
	OtpauthUrl string `json:"otpauth_url" valid:"required"`
}
