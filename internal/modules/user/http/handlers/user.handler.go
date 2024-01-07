package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/dto"
	"github.com/go-chi/jwtauth"
	"github.com/rs/zerolog/log"
	// "github.com/go-chi/chi"
)

type UserHandler struct {
	userService *services.UserService
	JwtExpiriesIn int
	ctx           context.Context
}

func NewUserHandler(userService *services.UserService, expiry int) *UserHandler {
	return &UserHandler{
		userService: userService,
		JwtExpiriesIn: expiry,
	}
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.UserInput
	var err error

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Err(err).Msg("Error decode request")
		return
	}
	err = u.userService.Create(user.Name, user.LastName, user.CPF, user.Email, user.Password, user.Admin)
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
		w.WriteHeader(http.StatusUnauthorized)
		log.Error().Err(err).Msg("")
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid or missing token"})
		return
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid token"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sub)
}

func (u *UserHandler) Authentication(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)
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
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// var user dto.AuthenticationInput
}
