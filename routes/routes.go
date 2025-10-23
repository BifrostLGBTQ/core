// routes/router.go
package routes

import (
	"bifrost/constants"
	"bifrost/handlers"
	"bifrost/router"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Router struct {
	mux    *mux.Router
	action *router.ActionRouter
}

func NewRouter(db *gorm.DB) *Router {
	r := &Router{
		mux:    mux.NewRouter(),
		action: router.NewActionRouter(),
	}

	// Action register
	r.action.Register(constants.CMD_INITIAL_SYNC, func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleInitialSync(db)(w, r)
	})
	r.action.Register(constants.CMD_AUTH_LOGIN, handlers.HandleLogin)
	// Diğer actionlar da buraya eklenebilir

	r.mux.HandleFunc("/", r.handlePacket)

	// Tek packet endpoint
	r.mux.HandleFunc("/packet", r.handlePacket)

	return r
}

func (r *Router) handlePacket(w http.ResponseWriter, req *http.Request) {
	// Parse GET ve POST verilerini birleştir
	if err := req.ParseForm(); err != nil {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		return
	}

	action := req.FormValue("action")

	if action == "" {
		// action yoksa default handler
		fmt.Println("Default handler çalıştı")
		w.Write([]byte("Default handler executed"))
		return
	}

	// action varsa ActionRouter kullan
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
