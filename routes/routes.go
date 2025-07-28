package routes

import (
	"bifrost/constants"
	"bifrost/services/db"
	"bifrost/services/socket"
	"fmt"
	"net/http"

	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type Router struct {
	// Burada router'a özgü alanlar ve bağımlılıklar yer alacak
	// Örnek:
	mux *mux.Router
}

// Router struct'ı, AppHandler arayüzünü uygulayacak şekilde ServeHTTP methodunu tanımlayalım
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	r.mux.ServeHTTP(w, req)
}

func NewRouter() *Router {
	go socket.ListenServer()
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path := "./static/" + req.URL.Path
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			http.NotFound(w, req)
			return
		} else if _, err := os.Stat(path); os.IsNotExist(err) {
			http.NotFound(w, req)
			return
		}
		fs.ServeHTTP(w, req)
	})))

	return &Router{
		mux: r,
	}
}

func handlePacket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("handlePacked:RECEIVE")
	err := r.ParseMultipartForm(250 << 20) // 250MB
	if err != nil {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		return
	}

	actionInt, err := strconv.Atoi(r.FormValue("action"))
	if err != nil {
		http.Error(w, constants.InvalidAction.String(), http.StatusBadRequest)
		return
	}

	userRepo := &db.UserRepositoryImpl{DB: db.DB}
	//	chatRepo := &db.ChatRepositoryImpl{DB: db.DB, SOCKET: socket.Server}

	action := constants.TCommandTypes(actionInt)

	// REGISTER
	if action == constants.ACT_AUTH_REGISTER {
		userRepo.Register(w, r.FormValue("email"), r.FormValue("password"), r.FormValue("captcha"))
		return
	}

	//LOGIN
	if action == constants.ACT_AUTH_LOGIN {
		userRepo.Login(w, r.FormValue("email"), r.FormValue("password"))
		return
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized 0", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader
		if strings.Contains(authHeader, " ") {
			tokenString = strings.Split(authHeader, " ")[1]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("USER_JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token geçerliyse sonraki işlemi çağırın
		next.ServeHTTP(w, r)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// Bu fonksiyon ile Router'a bir route ve handler ekleyebilirsiniz
func (r *Router) AddRoute(route string, handler http.Handler) {
	r.mux.Handle(route, handler)
}

// Bu fonksiyon ile Router'a bir route ve handlerFunc ekleyebilirsiniz
func (r *Router) AddHandlerFunc(route string, handlerFunc http.HandlerFunc) {
	r.mux.HandleFunc(route, handlerFunc)
}
