package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var u User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		rt.baseLogger.WithError(err).Error("Login request failed to parse body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u_db, err := rt.db.InitSetUserID(u.userToDatabase())
	u.userFromDatabase(u_db)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content_type", "application/json")
	_ = json.NewEncoder(w).Encode(u)
}

func (rt *_router) setUserID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var u User

	user_id := ps.ByName("user_id")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Request failed to parse user_name")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u_db, err := rt.db.SetUserID(u.userToDatabase(), user_id)
	u.userFromDatabase(u_db)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(u)
}

func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var u User

	user_name := ps.ByName("user_name")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Request failed to parse user_name")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u_db, err := rt.db.SetUsername(u.userToDatabase(), user_name)

	u.userFromDatabase(u_db)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(u)
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (rt *_router) getUserStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
