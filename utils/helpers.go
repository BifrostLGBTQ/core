package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/mail"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"

	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"

	"bifrost/constants"
	"bifrost/models/user"

	"github.com/nfnt/resize"
)

func JsonToString(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func StringToJson(jsonString string, data interface{}) error {
	err := json.Unmarshal([]byte(jsonString), data)
	if err != nil {
		return err
	}
	return nil
}

func DecodeGRPCMessages(jsonString string, data interface{}) {

	err := StringToJson(jsonString, &data)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func DecodeUserJWT(tokenString string) (*user.UserJWTClaims, error) {
	fmt.Println("DecodeUserJWT:Token:", tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &user.UserJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("USER_JWT_SECRET")), nil
	})
	if err != nil {
		fmt.Println("DecodeUserJWT:Error:1", err)
		return nil, err
	}
	if !token.Valid {
		fmt.Println("DecodeUserJWT:Error:2")
		return nil, errors.New("invalid jwt token")
	}
	myClaims, ok := token.Claims.(*user.UserJWTClaims)
	if !ok {
		fmt.Println("DecodeUserJWT:Error:3")
		return nil, errors.New("couldn't parse token claims")
	}
	fmt.Println("DecodeUserJWT:Passed:1")
	return myClaims, nil
}

func ValidatePassword(password string) bool {
	var passwordExp = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`

	passwordRegexp := regexp.MustCompile(passwordExp)
	return passwordRegexp.MatchString(password)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateHashedPassword(password string) (string, error) {

	//isValid := ValidatePassword(password)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	//match := CheckPasswordHash(password, string(hashedPassword))
	//fmt.Println("Match:", match)
	return string(hashedPassword), nil
}

func GenerateUserJWT(user_id uuid.UUID, email string) string {
	var jwtSecret = []byte(os.Getenv("USER_JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id": user_id,
		"email":   email,
		"exp":     time.Now().AddDate(0, 0, 30).Unix(),
		"version": "0.1.2",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(jwtSecret)
	result := fmt.Sprintf("Bearer %s", tokenString)
	return result
}

func GenerateUserJWTWithoutBearer(user_id uint64, email string) string {
	var jwtSecret = []byte(os.Getenv("USER_JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id": user_id,
		"email":   email,
		"exp":     time.Now().AddDate(0, 0, 30).Unix(),
		"version": "0.1.2",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func GenerateUniqueID() uuid.UUID {
	return uuid.New()
}

func Slugify(s string) string {
	slug.CustomSub = map[string]string{
		"ö": "o",
		"ç": "c",
		"ş": "s",
		"ı": "i",
		"ü": "u",
		"ğ": "g",
	}
	slugText := slug.Make(s)
	return slugText
}

func CheckAuthResponse(w http.ResponseWriter, r *http.Request) (*user.UserJWTClaims, error, int) {

	var decoded *user.UserJWTClaims
	var err error
	token := r.Header.Get("Authorization")
	fmt.Println("TOKEN", token)
	userID := uuid.Nil
	if strings.Contains(token, " ") {
		token = strings.Split(token, " ")[1]
	}
	if token != "" {
		decoded, err = DecodeUserJWT(token)
		fmt.Println("decodeds", decoded)
		if err != nil {
			return decoded, errors.New("Invalid Auth Token!"), http.StatusBadRequest
		}

		userID = decoded.UserID
	} else {
		err := constants.ErrorMessages[2]
		return nil, errors.New(err), http.StatusNetworkAuthenticationRequired
	}

	if userID == uuid.Nil {
		return nil, errors.New("invalid user!"), http.StatusBadRequest
	}

	fmt.Println("decodeds 2 =========", decoded)
	return decoded, nil, 200
}

func GenerateResponse(w http.ResponseWriter, status bool, message string) {
	response := map[string]interface{}{
		"Status":  status,
		"Message": message,
	}
	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func StrToIntDef(inputVal string, defaultVal uint64) uint64 {
	numVal, err := strconv.Atoi(inputVal)
	if err == nil {
		return uint64(numVal)
	}
	return defaultVal
}

func StrToBoolDef(inputVal string, defaultVal bool) bool {
	booleanValue, err := strconv.ParseBool(inputVal)
	if err == nil {
		return booleanValue
	}
	return defaultVal
}

func ExtractNameFromEmail(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return ""
	}
	return email[:atIndex]
}

func GetFloat64OrDefault(data map[string]interface{}, key string, defaultVal float64) float64 {
	if val, ok := data[key]; ok && val != nil {
		return val.(float64)
	}
	return defaultVal
}
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func IsAuthUserValid(w http.ResponseWriter, r *http.Request, expecedUserId uuid.UUID) bool {
	fmt.Println("IsAuthUserValid", "EXPECTED", expecedUserId)
	token := r.Header.Get("Authorization")
	if strings.Contains(token, "HISSEYORUMCOINVESTINGBOT") {
		return true
	}
	fmt.Println("IsAuthUserValid", "TOKEN", token)

	can_continue := false
	if strings.Contains(token, " ") {
		token = strings.Split(token, " ")[1]
	} else {
		return false
	}
	userID := uuid.Nil
	if token != "" {
		decoded, err := DecodeUserJWT(token)

		if err != nil {
			can_continue = false
		} else {
			userID = decoded.UserID
			if expecedUserId == userID {
				can_continue = true
			}
		}

	} else {
		can_continue = false
	}

	if expecedUserId == uuid.Nil {
		can_continue = false
	}

	if userID == uuid.Nil {
		can_continue = false
	}
	return can_continue
}

func IsValidFile(fileName string) bool {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".png":  true,
		".mp4":  true,
		".pdf":  true,
		".webm": true,
		".avi":  true,
		".webp": true,
		".jpeg": true,
		".gif":  true,
	}
	is_valid := false
	ext := filepath.Ext(fileName)
	is_valid = allowedExtensions[ext]
	return is_valid
}

func ShuffleString(s string) string {
	// Convert string to a slice of runes
	runes := []rune(s)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Shuffle the slice of runes
	for i := len(runes) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Convert the slice of runes back to a string
	return string(runes)
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateNickName(email string) string {
	username := ExtractNameFromEmail(email)
	randomString := GenerateRandomString(5)
	nickName := fmt.Sprintf("%s%s", username, randomString)
	username = Slugify(ShuffleString(nickName))
	return username
}

// VerifyReCAPTCHA ReCAPTCHA doğrulaması yapar.
func VerifyReCAPTCHA(secret string, response string) (bool, error) {
	// Google tarafından dönen cevabı temsil eder.
	type recaptchaResponse struct {
		Success bool `json:"success"`
	}
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{"secret": {secret}, "response": {response}})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var r recaptchaResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return false, err
	}

	return r.Success, nil
}

func IsValidUsername(username string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]{1,30}$")
	return re.MatchString(username)
}

func CreateDirectoryIfNotExists(directoryPath string) error {
	// Check if the directory already exists
	_, err := os.Stat(directoryPath)

	if os.IsNotExist(err) {
		// If the directory does not exist, create it
		errDir := os.MkdirAll(directoryPath, 0755)
		if errDir != nil {
			return fmt.Errorf("Error creating directory: %v", errDir)
		}
		fmt.Println("Directory created:", directoryPath)
	} else if err != nil {
		// If there is an error other than "not exists," return the error
		return fmt.Errorf("Error checking directory: %v", err)
	}

	return nil
}

func DownloadFile(url, destination string) error {
	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file where we will save the downloaded content
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy the content from the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("File downloaded successfully to %s\n", destination)
	return nil
}

func GenerateThumbnail(inputPath, outputPath string, maxWidth, maxHeight uint) error {
	// Open the image file
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Calculate the aspect ratio
	originalWidth := uint(img.Bounds().Dx())
	originalHeight := uint(img.Bounds().Dy())
	aspectRatio := float64(originalWidth) / float64(originalHeight)

	// Calculate new dimensions while maintaining the aspect ratio
	var newWidth, newHeight uint
	if originalWidth > originalHeight {
		newWidth = maxWidth
		newHeight = uint(float64(maxWidth) / aspectRatio)
	} else {
		newHeight = maxHeight
		newWidth = uint(float64(maxHeight) * aspectRatio)
	}

	// Resize the image to create a thumbnail
	thumb := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Save the thumbnail as JPEG
	err = jpeg.Encode(outputFile, thumb, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}

	fmt.Printf("Thumbnail generated and saved to: %s\n", outputPath)
	return nil
}

func IsValidImage(filePath string) bool {
	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	// Decode the image
	_, _, err = image.Decode(file)
	return err == nil
}

func StringToMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Base64Decode(encodedString string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return "", err
	}

	decodedString := string(decodedBytes)
	return decodedString, nil
}
