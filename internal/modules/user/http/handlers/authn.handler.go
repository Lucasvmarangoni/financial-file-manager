package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Lucasvmarangoni/logella/err"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/dto"
	"github.com/go-chi/jwtauth"
	"github.com/rs/zerolog/log"
)

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.UserInput  true  "user data"
// @Success      200
// @Failure      400
// @Router       /authn/create [post]
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.UserInput
	var wg sync.WaitGroup

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Err(err).Msg("Error decode request")
		return
	}

	// _, err = govalidator.ValidateStruct(user)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	log.Error().Err(err).Msg("Validation failed")
	// 	return
	// }

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = u.userService.Create(user.Name, user.LastName, user.CPF, user.Email, user.Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error().Stack().Err(err).Msg("Error create user ")
			return
		}				
	}()
	wg.Wait()
	w.WriteHeader(http.StatusOK)
}

// Authentication godoc
// @Summary      Generate a user JWT
// @Description  Generate a user JWT
// @Tags         Authn
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      400
// @Failure      401
// @Router       /authn/jwt [post]
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

func (u *UserHandler) GetSub(w http.ResponseWriter, r *http.Request) (string, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return "", errors.ErrCtx(err, "Failed to get JWT claims")
	}
	id, ok := claims["sub"].(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return "", errors.ErrCtx(err, "sub claim is missing or not a string")
	}
	return id, nil
}
