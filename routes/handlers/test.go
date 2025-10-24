package handlers

import (
	"bifrost/constants"
	services "bifrost/services/user"
	"bifrost/utils"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type TestHandler struct {
	service *services.UserService
}

func NewTestHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func HandleTestUser(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		s.Test()

		userID := "ce707772-f946-4ce3-9f89-e1fc091985f4"

		uid, err := uuid.Parse(userID)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		fmt.Println("TEST1", userID)
		userObj, err := s.GetUserByID(uid)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, constants.ErrDatabaseError)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user": userObj,
		})
	}
}
