package entities

import (
	"errors"
	"regexp"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        entities.ID `json:"id" valid:"required"`
	Name      string      `json:"name" valid:"length(3|10),alpha"`
	LastName  string      `json:"last_name" valid:"length(3|50),alpha"`
	CPF       string      `json:"cpf" valid:"matches(^[0-9]{3}\\.[0-9]{3}\\.[0-9]{3}-[0-9]{2}$)"`
	Email     string      `json:"email" valid:"email"`
	Password  string      `json:"password" valid:"required"`
	Admin     bool        `json:"admin" valid:"-"`
	CreatedAt time.Time   `json:"created_at" valid:"required"`
	UpdatedAt []time.Time `json:"updated_at" valid:"-"`
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ValidateHashPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func validatePassword(password string) error {
	if len(password) < 10 {
		return errors.New("password must be at least 10 characters long")
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("password must contain at least one digit")
	}

	if !regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func NewUser(name, lastName, cpf, email, password string, admin bool) (*User, error) {

	err := validatePassword(password)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Name:     name,
		LastName: lastName,
		CPF:      cpf,
		Email:    email,
		Password: string(hash),
		Admin:    admin,
	}
	user.prepare()

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) prepare() {
	if len(u.UpdatedAt) == 0 && u.CreatedAt.IsZero() {
		u.ID = entities.NewID()
		u.CreatedAt = time.Now()
	}
}

func (u *User) Update(id entities.ID, createdAt time.Time) {
	u.UpdatedAt = append(u.UpdatedAt, time.Now())
	u.CreatedAt = createdAt
	u.ID = id
}
