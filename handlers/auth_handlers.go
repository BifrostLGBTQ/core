package handlers

import (
	"fmt"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LOGIN HANDLED ✅")
}
