package handlers

import (
	"bifrost/middleware"
	services "bifrost/services/user"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
)

type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func HandleCreate(s *services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(5 * 1024 * 1024 * 1024)
		if err != nil {
			http.Error(w, "Could not parse multipart form: "+err.Error(), http.StatusBadRequest)
			return
		}

		formParams := r.MultipartForm.Value        // text fields
		images := r.MultipartForm.File["images[]"] // images array
		videos := r.MultipartForm.File["videos[]"] // videos array

		files := append([]*multipart.FileHeader{}, images...)
		files = append(files, videos...)

		fmt.Println("FILES", files, images, videos)

		user, ok := middleware.GetAuthenticatedUser(r)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}
		post, err := s.CreatePost(formParams, files, user)
		if err != nil {
			http.Error(w, "Failed to create post: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)

	}
}

func HandleGetByID(s *services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing post id", http.StatusBadRequest)
			return
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid uuid", http.StatusBadRequest)
			return
		}

		post, err := s.GetPostByID(id)
		if err != nil {
			http.Error(w, "failed to get post: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}
