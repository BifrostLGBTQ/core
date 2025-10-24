// routes/router.go
package routes

import (
	"bifrost/constants"
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

	// repository ve service oluştur
	userRepo := repositories.NewUserRepository(r.db)
	userService := services.NewUserService(userRepo)
	// Action register
	r.action.Register(constants.CMD_INITIAL_SYNC, handlers.HandleInitialSync(r.db))
	r.action.Register(constants.CMD_AUTH_REGISTER, handlers.HandleRegister(userService))
	r.action.Register(constants.CMD_AUTH_LOGIN, handlers.HandleLogin(userService))
	r.action.Register(constants.CMD_AUTH_TEST, handlers.HandleTestUser(userService))

	r.mux.HandleFunc("/", r.handlePacket)
	r.mux.HandleFunc("/test", r.handlePacket)

	// Tek packet endpoint
	r.mux.HandleFunc("/packet", r.handlePacket)

	return r
}
func (r *Router) handlePacket(w http.ResponseWriter, req *http.Request) {
	var action string

	if req.Method == http.MethodGet {
		// GET query parametrelerinden al
		action = req.URL.Query().Get("action")
	} else if req.Method == http.MethodPost {
		// POST ise önce JSON deneyebiliriz
		if strings.Contains(req.Header.Get("Content-Type"), "application/json") {
			var packet struct {
				Action string `json:"action"`
			}
			if err := json.NewDecoder(req.Body).Decode(&packet); err != nil {
				http.Error(w, "invalid JSON body", http.StatusBadRequest)
				return
			}
			action = packet.Action
		} else {
			// POST form / multipart
			if err := req.ParseMultipartForm(8192 << 20); err != nil {
				http.Error(w, "Could not parse form", http.StatusBadRequest)
				return
			}
			action = req.FormValue("action")
		}
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if action == "" {
		fmt.Println("Default handler çalıştı")
		w.Write([]byte("Default handler executed"))
		return
	}

	handler, ok := r.action.GetHandler(action)
	if !ok {
		http.Error(w, "Unknown action", http.StatusBadRequest)
		return
	}

	handler(w, req)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
