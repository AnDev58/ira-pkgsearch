// Package auth is a model specifying user database
package auth

import (
	"crypto/sha256"
)

// User is a model for user accounts
type User struct {
	Username string // Username is a user's login
	password [32]byte
	uid      int64
	gid      int64
	info     []byte
}

// CreateUser creates new User from username, password and other info
func CreateUser(username string, password string, info []byte) *User {
	var user User
	user.Username = username
	user.password = sha256.Sum256([]byte(password))
	user.info = info
	return &user
}
