package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// User represents a user entity.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handleGetAll(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement database query
	users := []User{
		{ID: "1", Name: "Max Mustermann", Email: "max@example.com"},
		{ID: "2", Name: "Erika Muster", Email: "erika@example.com"},
	}

	respondJSON(w, http.StatusOK, users)
}

func handleGetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// TODO: Implement database query
	user := User{ID: id, Name: "Max Mustermann", Email: "max@example.com"}

	respondJSON(w, http.StatusOK, user)
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// TODO: Implement database insert
	user.ID = "3" // Mock ID

	respondJSON(w, http.StatusCreated, user)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// TODO: Implement database update
	user.ID = id

	respondJSON(w, http.StatusOK, user)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")

	// TODO: Implement database delete

	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
