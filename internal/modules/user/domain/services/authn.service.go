package services

import (
	"strings"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/go-chi/jwtauth"
)

func (u *UserService) Authn(unique, password string, jwt *jwtauth.JWTAuth, jwtExpiresIn int) (string, error) {
	var user *entities.User
	var err error
	var operation string

	if strings.Contains(unique, "@") {
		operation = "FindByEmail"
		user, err = u.FindByEmail(unique, nil)
	} else {
		operation = "FindByCpf"
		user, err = u.FindByCpf(unique, nil)
	}
	if err != nil {
		return "", errors.NewError(err, operation)
	}

	err = user.ValidateHashPassword(password)
	if err != nil {
		return "", errors.NewError(err, "user.ValidateHashPassword")
	}

	user.Password = ""

	tokenString, err := u.GenerateJWT(user, jwt, jwtExpiresIn)
	if err != nil {
		return "", errors.NewError(err, "u.GenerateJWT")
	}
	return tokenString, nil
}

func (u *UserService) GenerateJWT(user *entities.User, jwt *jwtauth.JWTAuth, jwtExpiresIn int) (string, error) {
	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	if err != nil {
		return "", errors.NewError(err, "jwt.Encode")
	}
	return tokenString, nil
}
