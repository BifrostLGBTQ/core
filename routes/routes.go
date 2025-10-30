// routes/router.go
package routes

import (
	"bifrost/constants"
	"bifrost/middleware"
	"bifrost/repositories"
	"bifrost/router"
	"bifrost/routes/handlers"
	services "bifrost/services/user"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Router struct {
	mux    *mux.Router
	action *router.ActionRouter
	db     *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	r := &Router{
		mux:    mux.NewRouter(),
		action: router.NewActionRouter(db),
		db:     db,
	}

	r.mux.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static")),
		))

	// repository ve service oluştur
	userRepo := repositories.NewUserRepository(r.db)
	mediaRepo := repositories.NewMediaRepository(r.db)
	postRepo := repositories.NewPostRepository(r.db)

	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(userRepo, postRepo, mediaRepo)

	r.action.Register(constants.CMD_INITIAL_SYNC, handlers.HandleInitialSync(r.db)) // middleware yok

	r.action.Register(constants.CMD_INITIAL_SYNC, handlers.HandleInitialSync(r.db))

	// Action register
	r.action.Register(constants.CMD_AUTH_REGISTER, handlers.HandleRegister(userService))
	r.action.Register(constants.CMD_AUTH_LOGIN, handlers.HandleLogin(userService))
	r.action.Register(constants.CMD_AUTH_TEST, handlers.HandleTestUser(userService))

	r.action.Register(
		constants.CMD_USER_UPLOAD_AVATAR,
		handlers.HandleUploadAvatar(userService), // handler
		middleware.AuthMiddleware(userRepo),      // middleware
	)

	r.action.Register(
		constants.CMD_USER_UPLOAD_COVER,
		handlers.HandleUploadCover(userService), // handler
		middleware.AuthMiddleware(userRepo),     // middleware
	)

	r.action.Register(
		constants.CMD_USER_UPLOAD_STORY,
		handlers.HandleUploadStory(userService), // handler
		middleware.AuthMiddleware(userRepo),     // middleware
	)

	// POST
	//	r.action.Register(constants.CMD_POST_CREATE, middleware.AuthMiddleware(userRepo) handlers.HandleCreate(postService))
	r.action.Register(
		constants.CMD_POST_CREATE,
		handlers.HandleCreate(postService),  // handler
		middleware.AuthMiddleware(userRepo), // middleware
	)
	r.action.Register(constants.CMD_POST_FETCH, handlers.HandleGetByID(postService))
	r.action.Register(constants.CMD_POST_TIMELINE, handlers.HandleTimeline(postService))

	r.mux.HandleFunc("/", r.handlePacket)
	r.mux.HandleFunc("/test", r.handlePacket)

	// Tek packet endpoint
	r.mux.HandleFunc("/packet", r.handlePacket)

	return r
}

func (r *Router) handlePacket(w http.ResponseWriter, req *http.Request) {
	var action string

	switch req.Method {
	case http.MethodGet:
		// GET query parametrelerinden al
		action = req.URL.Query().Get("action")

	case http.MethodPost:
		contentType := req.Header.Get("Content-Type")
		if strings.Contains(contentType, "application/json") {
			// JSON body
			var packet struct {
				Action string `json:"action"`
			}
			if err := json.NewDecoder(req.Body).Decode(&packet); err != nil {
				http.Error(w, "invalid JSON body", http.StatusBadRequest)
				return
			}
			action = packet.Action
		} else {
			// Form / multipart
			if err := req.ParseMultipartForm(8192 << 20); err != nil {
				http.Error(w, "Could not parse form", http.StatusBadRequest)
				return
			}
			action = req.FormValue("action")
		}

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if action == "" {
		fmt.Println("Default handler çalıştı")
		w.Write([]byte("Default handler executed"))
		return
	}

	route, ok := r.action.GetHandler(action)
	if !ok {
		http.Error(w, "Unknown action", http.StatusBadRequest)
		return
	}

	// Middleware zincirini uygula
	handler := route.Handler
	for i := len(route.Middlewares) - 1; i >= 0; i-- {
		handler = route.Middlewares[i](handler)
	}

	// Handler çalıştır
	handler.ServeHTTP(w, req)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
