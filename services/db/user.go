package db

import (
	"bifrost/models"
	"bifrost/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	Register(user *models.User)
	Login(user *models.User)
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *UserRepositoryImpl) CreateUser(user *models.User) (*models.User, error) {
	result := repo.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUser(condition *models.User) (models.User, error) {
	var user models.User
	err := repo.DB.Preload("Followers").Preload("Followings").Preload("BlockedUsers").Preload("FollowedEquities").Where(condition).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (repo *UserRepositoryImpl) UpdateUser(condition models.User, updateFields map[string]interface{}) error {
	if err := repo.DB.Model(&models.User{}).Where(condition).Updates(updateFields).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepositoryImpl) Register(w http.ResponseWriter, email, password, captcha string) {

	if password == "" {
		utils.GenerateResponse(w, false, "Geçersiz şifre!")
		return
	}

	if email == "" {
		utils.GenerateResponse(w, false, "Geçersiz email!")
		return
	}

	if captcha == "" {
		utils.GenerateResponse(w, false, "Geçersiz captcha doğrulama")
		return
	}

	captchaValid, captchaErr := utils.VerifyReCAPTCHA("6Lc011MoAAAAAIHA0U6jGQX9198p1VWteUsfyCzN", captcha)
	if captchaErr != nil {
		utils.GenerateResponse(w, false, "Geçersiz captcha doğrulama kodu")
		return
	}

	if !captchaValid {
		utils.GenerateResponse(w, false, "Geçersiz captcha doğrulama başarılı olamadı!")
		return
	}

	isValidEmail := utils.ValidateEmail(email)

	if !isValidEmail {
		utils.GenerateResponse(w, false, "Email adresi doğrulanamadı!")
		return
	}
	validate := validator.New()

	user := &models.User{Email: email, Password: password}

	if err := validate.Struct(user); err != nil {
		validationErrors := make([]string, 0)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("Invalid %s field", fieldErr.Field()))
		}
		http.Error(w, fmt.Sprintf("%v", validationErrors), http.StatusBadRequest)
		return
	}

	hashed_pw, err := utils.GenerateHashedPassword(password)
	if err != nil {
		http.Error(w, "Password Error", http.StatusBadRequest)
	}

	username := utils.GenerateNickName(email)
	newUser := &models.User{
		Email:    email,
		Password: hashed_pw,
		UserName: username,
	}
	user, err = repo.CreateUser(newUser)

	if err != nil {
		utils.GenerateResponse(w, false, "Email adresi zaten sisteme kayıtlı!")
		return
	}

	token := utils.GenerateUserJWT(user.ID, user.Email)

	response := map[string]interface{}{
		"Status": true,
		"Token":  token,
		"User":   user,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (repo *UserRepositoryImpl) Login(w http.ResponseWriter, email, password string) {

	if password == "" {
		utils.GenerateResponse(w, false, "Şifre boş olamaz!")
		return
	}

	if email == "" {
		utils.GenerateResponse(w, false, "Email adresi boş olamaz!")
		return
	}

	validate := validator.New()
	user := models.User{Email: email, Password: password}
	if err := validate.Struct(user); err != nil {
		validationErrors := make([]string, 0)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("Invalid %s field", fieldErr.Field()))
		}
		http.Error(w, fmt.Sprintf("%v", validationErrors), http.StatusBadRequest)
		return
	}
	fields := models.User{Email: email}
	user, err := repo.GetUser(&fields)

	if err != nil {
		utils.GenerateResponse(w, false, "Kayıtlı kullanıcı bulunamadı!")
		return
	}

	isValid := utils.CheckPasswordHash(password, user.Password)

	if !isValid {
		utils.GenerateResponse(w, false, "Girilen şifre geçersiz!")
		return
	}

	token := utils.GenerateUserJWT(user.ID, email)

	response := map[string]interface{}{
		"Status": true,
		"Token":  token,
		"User":   user,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
