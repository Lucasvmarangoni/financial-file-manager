package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/user/http/dto"
	"github.com/go-chi/jwtauth"
	"github.com/rs/zerolog/log"
)

// Get user godoc
// @Summary      Get me user data
// @Description  Get me user data
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200 		 {object}  dto.UserOutput
// @Failure      500  
// @Router       /user/me [get]
// @Security ApiKeyAuth
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
		CreatedAt: finduser.CreatedAt,
		UpdateLog: finduser.UpdateLog,
	}
	userJSON, err := json.MarshalIndent(user, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

// Update user godoc
// @Summary      Update user
// @Description  Update user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.UserInput  true  "user data update"
// @Success      200
// @Failure      400  
// @Router       /user/update [put]
// @Security ApiKeyAuth
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
	log.Info().Str("context", "UserHandler").Msgf("User updated successfully (%s)", id)
	w.WriteHeader(http.StatusOK)
}

// Delete user godoc
// @Summary      Delete user
// @Description  Delete user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400 
// @Router       /user/del [delete]
// @Security ApiKeyAuth
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
	log.Info().Str("context", "UserHandler").Msgf("User deleted successfully (%s)", id)
	w.WriteHeader(http.StatusOK)
}
