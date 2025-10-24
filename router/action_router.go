// router/action_router.go
package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

// ActionRouter maps action strings to handler functions
type ActionRouter struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
	db       *gorm.DB
}

// NewActionRouter initializes a new ActionRouter
func NewActionRouter(db *gorm.DB) *ActionRouter {
	return &ActionRouter{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
		db:       db,
	}
}

// Register assigns a handler to an action string
func (ar *ActionRouter) Register(action string, handler func(http.ResponseWriter, *http.Request)) {
	ar.handlers[action] = handler
}

// RegisterDefault sets a handler that will be used when no action is provided
func (ar *ActionRouter) RegisterDefault(handler func(http.ResponseWriter, *http.Request)) {
	ar.handlers["__default__"] = handler
}

// Resolve finds and executes the handler for a given action
func (ar *ActionRouter) Resolve(w http.ResponseWriter, r *http.Request) {
	// Önce form ve query parametrelerini parse et
	r.ParseMultipartForm(10 << 20) // 10MB
	action := r.FormValue("action")
	if action == "" {
		action = r.URL.Query().Get("action")
	}

	// Eğer hala boşsa JSON body kontrol et
	if action == "" && r.Header.Get("Content-Type") == "application/json" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil && len(body) > 0 {
			var data map[string]interface{}
			if json.Unmarshal(body, &data) == nil {
				if a, ok := data["action"].(string); ok {
					action = a
				}
			}
		}
	}

	// Eğer hala action yoksa default handler çalışsın
	if action == "" {
		if handler, ok := ar.handlers["__default__"]; ok {
			handler(w, r)
			return
		}
		http.Error(w, "Missing action", http.StatusBadRequest)
		return
	}

	// Action varsa ilgili handler çağrılır
	handler, ok := ar.handlers[action]
	if !ok {
		http.Error(w, "Unknown action", http.StatusBadRequest)
		return
	}

	handler(w, r)
}

// GetHandler returns the handler for an action without executing it
func (ar *ActionRouter) GetHandler(action string) (http.HandlerFunc, bool) {
	h, ok := ar.handlers[action]
	return h, ok
}

func (ar *ActionRouter) GetHandlers() map[string]func(http.ResponseWriter, *http.Request) {
	return ar.handlers
}
