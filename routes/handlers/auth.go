package handlers

import (
	"bifrost/constants"
	"bifrost/middleware"
	"bifrost/models/user/payloads"
	services "bifrost/services/user"
	"bifrost/utils"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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
		fmt.Println("REGISTER:FORM", form)
		userObj, token, err := s.Register(form)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrUserExists)
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

		form := r.MultipartForm.Value

		userObj, token, err := s.Login(form)
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrInvalidInput)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user":  userObj,
			"token": token,
		})
	}
}

func HandleUploadAvatar(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		file, _, err := r.FormFile("avatar")
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}
		defer file.Close()

		fileHeader := r.MultipartForm.File["avatar"][0]

		newAvatar, err := s.UpdateAvatar(fileHeader, user)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, constants.ErrMediaUploadFailed)
			return
		}
		user.AvatarID = &newAvatar.ID
		user.Avatar = newAvatar

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user": user,
		})

	}
}

func HandleUploadCover(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		file, _, err := r.FormFile("cover")
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}
		defer file.Close()

		fileHeader := r.MultipartForm.File["cover"][0]

		newCover, err := s.UpdateCover(fileHeader, user)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, constants.ErrMediaUploadFailed)
			return
		}
		user.AvatarID = &newCover.ID
		user.Avatar = newCover

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user": user,
		})

	}
}

func HandleUploadStory(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		file, _, err := r.FormFile("story")
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}
		defer file.Close()

		fileHeader := r.MultipartForm.File["story"][0]

		newStory, err := s.AddStory(fileHeader, user)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, constants.ErrMediaUploadFailed)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"story": newStory,
		})

	}
}

func HandleUserInfo(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth_user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		userInfo, err := s.GetUserByID(auth_user.ID)
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user": userInfo,
		})
	}
}

func HandleSetUserAttribute(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth_user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		// Form verisini parse et
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		form := r.MultipartForm.Value
		attrIDs, exists := form["attribute_id"]
		if !exists || len(attrIDs) == 0 {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		var notes *string
		if noteVals, ok := form["notes"]; ok && len(noteVals) > 0 {
			notes = &noteVals[0]
		}

		// Tek attribute_id al
		attributeID, err := uuid.Parse(attrIDs[0])
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		attribute, err := s.GetAttribute(attributeID)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)

		}

		attr := &payloads.UserAttribute{
			CategoryType: attribute.Category,
			UserID:       auth_user.ID,
			AttributeID:  attributeID,
			Notes:        notes,
		}

		err = s.UpsertUserAttribute(attr)
		if err != nil {
			fmt.Println("ERROR", err)
			utils.SendError(w, http.StatusInternalServerError, constants.ErrUnknown)
			return
		}

		userInfo, err := s.GetUserByID(auth_user.ID)
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user": userInfo,
		})
	}
}

func HandleSetUserInterests(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth_user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		// Form verisini parse et
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		form := r.MultipartForm.Value
		attrIDs, exists := form["interest_id"]
		if !exists || len(attrIDs) == 0 {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		var notes *string
		if noteVals, ok := form["notes"]; ok && len(noteVals) > 0 {
			notes = &noteVals[0]
		}

		// Tek attribute_id al
		interestId, err := uuid.Parse(attrIDs[0])
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		interest, err := s.GetInterestItem(interestId)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)

		}

		userInterest := &payloads.UserInterest{

			UserID:         auth_user.ID,
			InterestItemID: interest.ID,
			Notes:          notes,
		}

		err = s.UpsertUserInterest(userInterest)
		if err != nil {
			fmt.Println("ERROR", err)
			utils.SendError(w, http.StatusInternalServerError, constants.ErrUnknown)
			return
		}

		userInfo, err := s.GetUserByID(auth_user.ID)
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user": userInfo,
		})
	}
}

func HandleSetUserFantasies(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth_user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		// Form verisini parse et
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		form := r.MultipartForm.Value
		fantasyIdRaw, exists := form["fantasy_id"]
		if !exists || len(fantasyIdRaw) == 0 {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		// Tek attribute_id al
		fantasyId, err := uuid.Parse(fantasyIdRaw[0])
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		fantasyInfo, err := s.GetFantasy(fantasyId)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)

		}

		fantasy := &payloads.UserFantasy{
			FantasyID: fantasyInfo.ID,
			UserID:    auth_user.ID,
		}

		err = s.UpsertUserFantasy(fantasy)
		if err != nil {
			fmt.Println("ERROR", err)
			utils.SendError(w, http.StatusInternalServerError, constants.ErrUnknown)
			return
		}

		userInfo, err := s.GetUserByID(auth_user.ID)
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user": userInfo,
		})
	}
}

func HandleSetUserSexualIdentities(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth_user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		// Form verisini parse et
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		form := r.MultipartForm.Value

		genderIDs, hasGender := form["gender_identity_id"]
		sexualIDs, hasSexual := form["sexual_orientation_id"]
		sexRoleIDs, hasRole := form["sexual_role_id"]

		// Üçü de boşsa hata döndür
		if (!hasGender || len(genderIDs) == 0) &&
			(!hasSexual || len(sexualIDs) == 0) &&
			(!hasRole || len(sexRoleIDs) == 0) {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		//

		fmt.Println("SEX ROLE IDS", sexRoleIDs)

		// En az birisi dolu olmalı
		if len(genderIDs) == 0 && len(sexualIDs) == 0 && len(sexRoleIDs) == 0 {
			utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		err := s.UpsertUserSexualIdentify(auth_user.ID, genderIDs, sexualIDs, sexRoleIDs)
		if err != nil {

			utils.SendError(w, http.StatusInternalServerError, constants.ErrInvalidInput)
			return
		}

		userInfo, err := s.GetUserByID(auth_user.ID)
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, constants.ErrUnauthorized)
			return
		}

		utils.SendJSON(w, http.StatusOK, map[string]interface{}{
			"user": userInfo,
		})
	}
}

/*

func Follow(w http.ResponseWriter, r *http.Request) {
	form := r.MultipartForm.Value
	followerID := form["follower_id"]
	followeeID := form["followee_id"]

	if len(followerID) == 0 || len(followeeID) == 0 {
		utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	if err := h.service.Follow(followerID[0], followeeID[0]); err != nil {
		utils.SendError(w, http.StatusBadRequest, constants.ErrDatabaseError)
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]string{
		"message": "User followed successfully",
	})
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	form := r.MultipartForm.Value
	followerID := form["follower_id"]
	followeeID := form["followee_id"]

	if len(followerID) == 0 || len(followeeID) == 0 {
		utils.SendError(w, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	if err := h.service.Unfollow(followerID[0], followeeID[0]); err != nil {
		utils.SendError(w, http.StatusBadRequest, constants.ErrDatabaseError)
		return
	}

	utils.SendJSON(w, http.StatusOK, map[string]string{
		"message": "User unfollowed successfully",
	})
}

*/
