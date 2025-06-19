// users.go
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"webserver/database"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetUserByID fetches a single user by ID from the DB
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	// Validate UUID format
	if _, err := uuid.Parse(userID); err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user User
	err := database.Pool.QueryRow(context.Background(),
		"SELECT id, username, email FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser inserts a new user into the DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Optional: Validate newUser.Username and newUser.Email here
	err := database.Pool.QueryRow(context.Background(),
		"INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id",
		newUser.Username, newUser.Email).Scan(&newUser.ID)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	rows, err := database.Pool.Query(context.Background(), "SELECT id, username, email FROM users")
	if err != nil {
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
			http.Error(w, "Failed to read users", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
