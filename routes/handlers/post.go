package handlers

import (
	"bifrost/middleware"
	services "bifrost/services/user"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

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

func HandleGetByPublicID(s *services.PostService) http.HandlerFunc {
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

		fmt.Println("%d", id)

		post, err := s.GetPostByPublicID(12)
		if err != nil {
			http.Error(w, "failed to get post: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func HandleTimeline(s *services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Limit parametresi
		limitStr := r.URL.Query().Get("limit")
		limit := 10 // default
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
				limit = l
			}
		}

		// Cursor parametresi (PublicID)
		var cursor *int64
		cursorStr := r.URL.Query().Get("cursor")
		if cursorStr != "" {
			if c, err := strconv.ParseInt(cursorStr, 10, 64); err == nil {
				cursor = &c
			} else {
				http.Error(w, "invalid cursor", http.StatusBadRequest)
				return
			}
		}

		// Timeline verisini çek
		result, err := s.GetTimeline(limit, cursor)
		if err != nil {
			http.Error(w, "failed to get timeline: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// JSON olarak döndür
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func HandleGetPostsByUserID(s *services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing post id", http.StatusBadRequest)
			return
		}

		userId, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid uuid", http.StatusBadRequest)
			return
		}

		limitStr := r.URL.Query().Get("limit")
		limit := 10 // default
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
				limit = l
			}
		}

		// Cursor parametresi (PublicID)
		var cursor *int64
		cursorStr := r.URL.Query().Get("cursor")
		if cursorStr != "" {
			if c, err := strconv.ParseInt(cursorStr, 10, 64); err == nil {
				cursor = &c
			} else {
				http.Error(w, "invalid cursor", http.StatusBadRequest)
				return
			}
		}

		post, err := s.GetPostsByUserID(userId, limit, cursor)
		if err != nil {
			http.Error(w, "failed to get post: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func HandleGetRepliesByUserID(s *services.PostService) http.HandlerFunc {
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

func HandleGetAllMediasByUserID(s *services.PostService) http.HandlerFunc {
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

func HandleGetAllLikesByUserID(s *services.PostService) http.HandlerFunc {
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
