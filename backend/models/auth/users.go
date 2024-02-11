// Package auth is a model specifying user database
package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	UID       int32 `db:"uid"`
	GID       int32 `db:"gid"`
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time `db:"created_at"`
	Info      string
}

// User is a model for user accounts
type User struct {
	Username string // Username is a user's login
	password [32]byte
	UID      int32
	GID      int32
	Info     any
	Email    string
	db       *sqlx.DB
}

// Constant - do not change!!
var errNoData = errors.New("Not enough data")

// ErrNoData returns error for no data
func ErrNoData() error {
	return errNoData
}

const (
	NoUID      = -1 // NULL for UID
	NoGID      = -1 // NULL for GID
	NoUsername = "" // NULL for Username
	NoPassword = ""
	NoEmail    = "@noemail" // NULL for email
)

// CreateUser creates new User in database from username, password, email and other info
func CreateUser(username string, password string, email string, info any, db *sqlx.DB) (*User, error) {
	user := NewUser(username, password, NoUID, NoGID, email, info, db)
	err := user.Save()
	return user, err
}

// NewUser fills User struct with all data
// WARNING: password must be sha256.Sum256 of real password
// All parameters except db may be missed by following constants:
// username (user's login): NoUsername
// uid (user's UID): NoUID
// gid (user's GID): NoGID
// email (user's email): NoEmail
// password (user's password): NoPassword
// info (something else that _may be converted to JSON_): nil
func NewUser(username string, password string, uid, gid int32, email string, info any, db *sqlx.DB) *User {
	u := &User{
		Username: username,
		password: sha256.Sum256([]byte(password)),
		UID:      uid,
		GID:      gid,
		Info:     info,
		db:       db,
		Email:    email,
	}
	if password == NoPassword {
		u.password = [32]byte{}
	}
	return u
}

func QueryUserByLogin(login string, db *sqlx.DB) (*User, error) {
	var user UserDB
	err := db.Get(&user, "SELECT * FROM public.users WHERE username = $1", login)
	if err != nil {
		return nil, err
	}
	decoded, err := hex.DecodeString(user.Password)
	if err != nil {
		return nil, err
	}
	var info any
	err = json.Unmarshal([]byte(user.Info), &info)
	return &User{Username: user.Username, password: [32]byte(decoded), UID: user.UID, GID: user.GID, Info: info, db: db, Email: user.Email}, nil
}

// Query is a wrapper for db.Query which creates a db select and returns users were found error
func (u *User) Query() ([]string, error) {
	var users []string
	var err error
	if u.UID != NoUID {
		err = u.db.Select(&users, "SELECT username FROM public.users WHERE uid = $1", u.UID)

	} else if u.Username != NoUsername {
		err = u.db.Select(&users, "SELECT username FROM public.users WHERE username = $1", u.Username)

	} else if u.Email != NoEmail {
		err = u.db.Select(&users, "username FROM public.users WHERE email = $1", u.Email)

	} else if u.Info != nil {
		info, jsonErr := json.Marshal(u.Info)
		if err != nil {
			return nil, jsonErr
		}

		if u.GID != NoGID {
			err = u.db.Select(&users, "SELECT username FROM public.users WHERE info = $1 AND gid = $2", info, u.GID)
		} else {
			err = u.db.Select(&users, "SELECT username FROM public.users WHERE info = $1", info)
		}

	} else if u.GID != NoGID {
		err = u.db.Select(&users, "SELECT username FROM public.users WHERE gid = $1", u.GID)
	} else {
		return nil, ErrNoData()
	}

	if err != nil {
		return nil, fmt.Errorf("auth.Query: %w", err)
	}
	return users, err
}

// Save updates or creates user in database
func (u *User) Save() error {
	users, err := u.Query()
	if err != nil {
		return err
	}
	if len(users) == 0 {
		if u.Username == NoUsername || u.Email == NoEmail || u.password == [32]byte{} || u.Info == nil {
			return ErrNoData()
		}
		_, err = u.db.Exec("INSERT INTO public.users (username, password, email, info, gid) VALUES ($1, $2, $3, $4, 6)", u.Username, fmt.Sprintf("%x", u.password), u.Email, u.Info)
		return err
	}

	if u.UID == NoUID {
		return ErrNoData()
	}
	transaction, err := u.db.Begin()
	defer transaction.Rollback()
	if err != nil {
		return err
	}
	if u.Username != NoUsername {
		simUser, err := NewUser(u.Username, NoPassword, NoUID, NoGID, NoEmail, nil, u.db).Query()
		if err != nil {
			return err
		}
		if len(simUser) != 0 {
			return fmt.Errorf("User with this username exists")
		}

		_, err = transaction.Exec("UPDATE public.users SET username = $1 WHERE uid = $2", u.Username, u.UID)
		if err != nil {
			return err
		}
	}
	if u.Email != NoEmail {
		simUser, err := NewUser(NoUsername, NoPassword, NoUID, NoGID, u.Email, nil, u.db).Query()
		if err != nil {
			return err
		}
		if len(simUser) != 0 {
			return fmt.Errorf("User with this username exists")
		}

		_, err = transaction.Exec("UPDATE public.users SET email = $1 WHERE uid = $2", u.Email, u.UID)
		if err != nil {
			return err
		}
	}
	if u.Info != nil {
		data, err := json.Marshal(u.Info)
		if err != nil {
			return err
		}
		_, err = transaction.Exec("UPDATE public.users SET info = $1 WHERE uid = $2", data, u.UID)
		if err != nil {
			return err
		}
	}
	if u.password != [32]byte{} {
		_, err = transaction.Exec("UPDATE public.users SET password = $1 WHERE uid = $2", fmt.Sprintf("%x", u.password), u.UID)
		if err != nil {
			return err
		}
	}
	return transaction.Commit()
}
