package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/domain/services"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/dto"
	"github.com/go-chi/jwtauth"
	// "github.com/go-chi/chi"
)

type UserHandler struct {
	createService *services.CreateService
	Jwt           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewUserHandler(createService *services.CreateService, jwt *jwtauth.JWTAuth, expiry int) *UserHandler {
	return &UserHandler{
		createService: createService,
		Jwt:           jwt,
		JwtExpiriesIn: expiry,
	}
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.UserInput
	var err error

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = u.createService.Create(user.Name, user.LastName, user.CPF, user.Email, user.Password, user.Admin)
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

// func (u *UserHandler) Authentication(w http.ResponseWriter, r *http.Request) {
// 	var user dto.AuthenticationInput
// }

// func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
// }

// func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
// 	var user dto.AuthenticationInput
// }
