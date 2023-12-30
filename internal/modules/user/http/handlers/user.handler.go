package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/dto"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/infra/repositories"
	"github.com/go-chi/jwtauth"
	// "github.com/go-chi/chi"
)

type UserHandler struct {
	Repository    *repositories.UserRepositoryDb
	Jwt           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.UserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) Me(w http.ResponseWriter, r *http.Request) {

	// id := chi.URLParam(r, "x")
	w.WriteHeader(http.StatusOK)
}


func (u *UserHandler) Authentication(w http.ResponseWriter, r *http.Request) {
	var user dto.AuthenticationInput
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var user dto.AuthenticationInput
}