package handlers

import (
	"context"

	"encoding/json"
	"net/http"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/dto"
	"github.com/Lucasvmarangoni/financial-file-manager/pkg/errors"
	"github.com/go-chi/jwtauth"
	"github.com/rs/zerolog/log"
	// "github.com/go-chi/chi"
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

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.UserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Err(err).Msg("Error decode request")
		return
	}
	err = u.userService.Create(user.Name, user.LastName, user.CPF, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Stack().Err(err).Msg("Error create user ")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("Failed to get JWT claims")
		return
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msg("sub claim is missing or not a string")
		return
	}
	finduser, err := u.userService.FindById(sub, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Stack().Msg("User not found")
		return
	}
	user := dto.UserOutput{
		ID:        finduser.ID,
		Name:      finduser.Name,
		LastName:  finduser.LastName,
		CPF:       finduser.CPF,
		Email:     finduser.Email,
		Admin:     finduser.Admin,
		CreatedAt: finduser.CreatedAt,
		UpdatedAt: finduser.UpdatedAt,
	}
	userJSON, err := json.MarshalIndent(user, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func (u *UserHandler) Authentication(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExpiresIn").(int)
	var user dto.AuthenticationInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Err(err).Msg("Error decode request")
		return
	}
	if user.Email != "" && user.CPF != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	unique := user.Email + user.CPF
	tokenString, err := u.userService.Authn(unique, user.Password, jwt, jwtExpiresIn)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Error().Stack().Err(err).Msg("Error authenticate user")
		return
	}

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user dto.UserUpdateInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Err(err).Msg("Error decode request")
		return
	}

	id, err := u.GetSub(w, r)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Error get sub")
	}

	err = u.userService.Update(id, user.Name, user.LastName, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Stack().Err(err).Msg("Error update user ")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := u.GetSub(w, r)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Error get sub")
	}
	err = u.userService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Stack().Err(err).Msg("Error delete user ")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) GetSub(w http.ResponseWriter, r *http.Request) (string, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return "", errors.NewError(err, "Failed to get JWT claims")
	}
	id, ok := claims["sub"].(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return "", errors.NewError(err, "sub claim is missing or not a string")
	}
	return id, nil
}
