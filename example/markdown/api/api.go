package api

import (
	"net/http"
	"time"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

type UsersCollection []User

type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

func ListUsers(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func GetUser(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func AddUser(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func UpdateUser(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
