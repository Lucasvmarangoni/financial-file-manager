package handlers

import (
	"context"
	"encoding/json"
	go_err "errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/validate"
)

type UserHandler struct {
	userService *services.UserService
	ctx         context.Context
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) validatePassword(password string, w http.ResponseWriter) {
	err := validate.ValidatePassword(password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "BadRequest",
			"message": fmt.Sprintf("Valid password is required. %v", err),
		})
		return
	}
}

func (u *UserHandler) validateCPF(cpf *string) error {
	cpfRegex := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	if *cpf != "" && !cpfRegex.MatchString(*cpf) {
		return go_err.New("invalid CPF format")
	}
	return nil
}

func (u *UserHandler) validateEmail(email *string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if *email != "" && !emailRegex.MatchString(*email) {
		return go_err.New("invalid email format")
	}
	return nil
}

func (u *UserHandler) validateNameAndLastname(name, lastname *string) error {
	nameRegex := regexp.MustCompile(`^[a-zA-Z ]+$`)

	if *name != "" && !nameRegex.MatchString(*name) {
		return go_err.New("invalid name format")
	}
	if *lastname != "" && !nameRegex.MatchString(*lastname) {
		return go_err.New("invalid last name format")
	}
	return nil
}
