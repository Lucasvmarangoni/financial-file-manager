package entities

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
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
	ID        entities.ID `json:"id" valid:"required"`
	Name      string      `json:"name" valid:"length(3|10),matches(^[a-zA-Z ]+$)"`
	LastName  string      `json:"last_name" valid:"length(3|50),matches(^[a-zA-Z ]+$)"`
	CPF       string      `json:"cpf" valid:"matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	HashCPF   string      `json:"hash_cpf" valid:"-"`
	Email     string      `json:"email" valid:"email"`
	HashEmail string      `json:"hash_email" valid:"-"`
	Password  string      `json:"password" valid:"required"`
	CreatedAt time.Time   `json:"created_at" valid:"required"`
	UpdateLog []UpdateLog `json:"update_log" valid:"-"`
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

func Hash(str string) string {
	hash := sha256.Sum256([]byte(str))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func NewUser(name, lastName, cpf, email, password string) (*User, error) {

	err := validate.ValidatePassword(password)
	if err != nil {
		return nil, errors.ErrCtx(err, "validatePassword")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrCtx(err, "bcrypt.GenerateFromPassword")
	}

	user := User{
		Name:      name,
		LastName:  lastName,
		CPF:       cpf,
		HashCPF:   Hash(cpf),
		Email:     email,
		HashEmail: Hash(email),
		Password:  string(hashPassword),
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

func (u *User) Update(oldValues []UpdateLog, id entities.ID, createdAt time.Time) {
	u.UpdateLog = oldValues
	u.CreatedAt = createdAt
	u.ID = id
}
