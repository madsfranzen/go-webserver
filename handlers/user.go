// users.go
package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/madsfranzen/go-webserver/database"
	"github.com/madsfranzen/go-webserver/models"
)

// CreateUser inserts a new user into the DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Optional: Validate newUser.Username and newUser.Email here
	err := database.Pool.QueryRow(context.Background(),
		"INSERT INTO users (username, email, premium) VALUES ($1, $2, $3) RETURNING id",
		newUser.Username, newUser.Email, newUser.Premium).Scan(&newUser.ID)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	rows, err := database.Pool.Query(ctx, `SELECT id, username, email, premium FROM users`)
	if err != nil {
		if ctx.Err() != nil {
			slog.Error("DB query canceled or timed out: ", "error", ctx.Err())
		} else {
			slog.Error("Query failed: ", "error", err)
		}
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Premium); err != nil {
			slog.Error("Scan error: ", "error", err)
			http.Error(w, "Failed to read users", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		slog.Error("Failed to encode response: ", "error", err)
	}
}
