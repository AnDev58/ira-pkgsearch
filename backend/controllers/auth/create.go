package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andev58/ira-pkgsearch/backend/models/auth"
	"github.com/andev58/ira-pkgsearch/backend/util"
)

// CreateUserHandler is a RESTful wrapper for auth.CreateUser
func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type RequestData struct {
		Username string          `json:"login"`
		Password string          `json:"password"`
		Email    string          `json:"email"`
		Info     json.RawMessage `json:"etc"`
	}

	if util.EnforceJSON(w, r) {
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var reqUser RequestData
	if err := decoder.Decode(&reqUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := auth.NewUser(reqUser.Username, reqUser.Password, auth.NoUID, auth.NoGID, reqUser.Email, nil, s.db)
	sameUsers, err := u.Query()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(sameUsers) > 0 {
		http.Error(w, "User with this username/email exists", http.StatusBadRequest)
		return
	}
	u.Info = reqUser.Info
	err = u.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, "OK")
	}
}
