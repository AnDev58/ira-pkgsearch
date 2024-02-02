// Package auth is a model specifying user database
package auth

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
)

// User is a model for user accounts
type User struct {
	username string // Username is a user's login
	password [32]byte
	uid      int64
	gid      int64
	info     any
	email    string
	db       *sql.DB
}

// Constant - do not change!!
var errNoData = errors.New("Not enough data")

// ErrNoData returns error for no data
func ErrNoData() error {
	return errNoData
}

const (
	NoUID      = -1         // NULL for UID
	NoGID      = -1         // NULL for GID
	NoUsername = ""         // NULL for Username
	NoEmail    = "@noemail" // NULL for email
)

// CreateUser creates new User in database from username, password, email and other info
func CreateUser(username string, password string, email string, info []byte, db *sql.DB) {
	var user User
	user.username = username
	user.password = sha256.Sum256([]byte(password))
	user.info = info
	user.db = db
}

// NewUser fills User struct with all data
// WARNING: password must be sha256.Sum256 of real password
// All parameters except db may be missed by following constants:
// username (user's login): NoUsername
// uid (user's UID): NoUID
// gid (user's GID): NoGID
// email (user's email): NoEmail
// password (password's SHA-256): [32]byte{}
// info (something else that _may be converted to JSON_): nil
func NewUser(username string, password [32]byte, uid, gid int64, email string, info any, db *sql.DB) *User {
	return &User{
		username: username,
		password: password,
		uid:      uid,
		gid:      gid,
		info:     info,
		db:       db,
		email:    email,
	}
}

// Query is a wrapper for db.Query which creates a db query and returns Rows or error
func (u *User) Query() (*sql.Rows, error) {
	if u.uid != NoUID {
		return u.db.Query("SELECT * FROM public.users WHERE uid = $1", u.uid)
	}
	if u.username != NoUsername {
		return u.db.Query("SELECT * FROM public.users WHERE name = $1", u.username)
	}
	if u.email != NoEmail {
		return u.db.Query("SELECT * FROM public.users WHERE email = $1", u.email)
	}
	if u.info != nil {
		info, err := json.Marshal(u.info)
		if err != nil {
			return nil, err
		}
		if u.gid != NoGID {
			return u.db.Query("SELECT * FROM public.users WHERE info = $1 AND gid = $2", info, u.gid)
		}
		return u.db.Query("SELECT * FROM public.users WHERE info = $1", info)
	}
	return nil, ErrNoData()
}

// Save updates or creates user in database
func (u *User) Save() error {
	_, err := u.Query()
	if err == sql.ErrNoRows {
		if u.username == NoUsername || u.email == NoEmail || u.password == [32]byte{} || u.info == nil {
			return ErrNoData()
		}
		// TODO: Validation
		_, err = u.db.Exec("INSERT INTO public.users VALUES (NULL, $1, $2, $3, NULL, info, 6)")
		return err
	}
	if err != nil {
		return err
	}
	// TODO: Update user
	return nil
}
