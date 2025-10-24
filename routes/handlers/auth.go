package handlers

import (
	"bifrost/constants"
	services "bifrost/services/user"
	"bifrost/utils"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func HandleRegister(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		form := r.MultipartForm.Value
		userObj, token, err := s.Register(form)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrDatabaseError)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user":  userObj,
			"token": token,
		})
	}
}

func HandleLogin(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

	}
}
