package entities

import (
	"crypto/sha256"
	"encoding/hex"
	go_err "errors"
	"regexp"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
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
		errors.ErrCtx(err, "bcrypt.CompareHashAndPassword")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 10 {
		return errors.ErrCtx(go_err.New("password must be at least 10 characters long"), "validatePassword")
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.ErrCtx(go_err.New("password must contain at least one lowercase letter"), "validatePassword")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.ErrCtx(go_err.New("password must contain at least one uppercase letter"), "validatePassword")
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.ErrCtx(go_err.New("password must contain at least one digit"), "validatePassword")
	}

	if !regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password) {
		return errors.ErrCtx(go_err.New("password must contain at least one special character"), "validatePassword")
	}

	return nil
}

func Hash(str string) string {
	hash := sha256.Sum256([]byte(str))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func NewUser(name, lastName, cpf, email, password string) (*User, error) {

	err := validatePassword(password)
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
