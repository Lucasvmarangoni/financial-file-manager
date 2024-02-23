package services

import (
	"strings"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/entities"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/go-chi/jwtauth"
)

func (u *UserService) Authn(unique, password string, jwt *jwtauth.JWTAuth, jwtExpiresIn int) (string, error) {
	var user *entities.User
	var err error
	var operation string

	if strings.Contains(unique, "@") {
		operation = "FindByEmail"
		user, err = u.FindByEmail(entities.Hash(unique), nil)
	} else {
		operation = "FindByCpf"
		user, err = u.FindByCpf(entities.Hash(unique), nil)
	}
	if err != nil {
		return "", errors.ErrCtx(err, operation)
	}

	err = user.ValidateHashPassword(password)
	if err != nil {
		return "", errors.ErrCtx(err, "user.ValidateHashPassword")
	}
	user.Password = ""

	tokenString, err := u.generateJWT(user, jwt, jwtExpiresIn)
	if err != nil {
		errors.ErrCtx(err, "u.generateJWT")
	}
	return tokenString, nil
}

func (u *UserService) generateJWT(user *entities.User, jwt *jwtauth.JWTAuth, jwtExpiresIn int) (string, error) {
	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),		
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	if err != nil {
		errors.ErrCtx(err, "jwt.Encode")
	}
	return tokenString, nil
}


