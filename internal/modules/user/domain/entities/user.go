package entities

import (
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/config"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/security"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/validate"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type UpdateLog struct {
	Timestamp time.Time
	OldValues map[string]interface{}
}

type User struct {
	ID         entities.ID `json:"id" valid:"required"`
	Name       string      `json:"name" valid:"length(3|10),matches(^[a-zA-Z ]+$)"`
	LastName   string      `json:"last_name" valid:"length(3|50),matches(^[a-zA-Z ]+$)"`
	CPF        string      `json:"cpf" valid:"matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	HashCPF    string      `json:"hash_cpf" valid:"-"`
	Email      string      `json:"email" valid:"email"`
	HashEmail  string      `json:"hash_email" valid:"-"`
	Password   string      `json:"password" valid:"required"`
	OtpSecret  string      `json:"otp_secret" valid:"-"`
	OtpAuthUrl string      `json:"otp_auth_url" valid:"-"`
	OtpEnabled bool        `json:"otp_enabled" valid:"-"`
	CreatedAt  time.Time   `json:"created_at" valid:"required"`
	UpdateLog  []UpdateLog `json:"update_log" valid:"-"`
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return errors.ErrCtx(err, "govalidator.ValidateStruct")
	}
	return nil
}

func (u *User) ValidateHashPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.ErrCtx(err, "bcrypt.CompareHashAndPassword")
	}
	return nil
}

func NewUser(name, lastName, cpf, email, password string) (*User, error) {

	hashPassword, err := bcryptHash(password)
	if err != nil {
		return nil, errors.ErrCtx(err, "u.PreparePassword")
	}

	hmac_key := []byte(config.GetEnvString("security", "hmac_key"))

	user := User{
		Name:      name,
		LastName:  lastName,
		CPF:       cpf,
		HashCPF:   security.HmacHash(cpf, hmac_key),
		Email:     email,
		HashEmail: security.HmacHash(email, hmac_key),
		Password:  hashPassword,
	}
	user.prepare()

	err = user.Validate()
	if err != nil {
		return nil, errors.ErrCtx(err, "user.Validate")
	}
	return &user, nil
}

func (u *User) prepare() {
	if len(u.UpdateLog) == 0 && u.CreatedAt.IsZero() {
		u.ID = entities.NewID()
		u.CreatedAt = time.Now()
	}
}

func (u *User) PrepateTOTP(otpSecret, otpAuthUrl string) {
	u.OtpSecret = otpSecret
	u.OtpAuthUrl = otpAuthUrl
}

func (u *User) Update(oldValues []UpdateLog, name, lastName, email, password string) error {
	u.UpdateLog = oldValues
	u.Name = name
	u.LastName = lastName
	u.Email = email

	hashPassword, err := bcryptHash(password)
	if err != nil {
		return errors.ErrCtx(err, "u.PreparePassword")
	}
	u.Password = string(hashPassword)
	return nil
}

func bcryptHash(input string) (string, error) {
	err := validate.ValidatePassword(input)
	if err != nil {
		return "", errors.ErrCtx(err, "validatePassword")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.ErrCtx(err, "bcrypt.GenerateFromPassword")
	}
	return string(hashPassword), nil
}
