package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"homestack/internal/service"

	"github.com/go-chi/chi/v5"
)

// Example provides CRUD handlers for the example User resource.
type Example struct {
	svc *service.ExampleService
}

// NewExample creates a new Example handler.
func NewExample(svc *service.ExampleService) *Example {
	return &Example{svc: svc}
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ListUsers handles GET /api/users.
func (e *Example) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := e.svc.ListUsers()
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser handles POST /api/users.
func (e *Example) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.Email == "" {
		http.Error(w, `{"error":"name and email are required"}`, http.StatusBadRequest)
		return
	}

	user, err := e.svc.CreateUser(req.Name, req.Email)
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser handles DELETE /api/users/{id}.
func (e *Example) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	if err := e.svc.DeleteUser(uint(id)); err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
