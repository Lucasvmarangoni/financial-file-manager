package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)


// Authentication godoc
// @Summary      Generate a user JWT
// @Description  Generate a user JWT
// @Tags         Authn
// @Accept       json
// @Produce      json
// @Success      200  {object}  
// @Failure      400  {object}  Error
// @Router      /authz/{id} [post]
func (u *UserHandler) AdminAuthz(w http.ResponseWriter, r *http.Request) {
	adminID, err := u.GetSub(w, r)
	userIDForAdminStatusToggle := chi.URLParam(r, "id")	

	if err != nil {
		log.Error().Stack().Err(err).Msg("Error get sub")
	}
	err = u.userService.AuthzAdmin(adminID, userIDForAdminStatusToggle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Stack().Err(err).Msg("Error provided admin authorization")
		return
	}
	log.Info().Str("context", "UserHandler").Msgf("User (%s) is a new Admin", userIDForAdminStatusToggle)
	w.WriteHeader(http.StatusOK)
}
