package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"server/db"
	"server/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type UserHandler struct {
	ctx context.Context
	l   *log.Logger
	q   *db.Queries
}

func NewUserHandler(ctx context.Context, l *log.Logger, pg *sql.DB) *UserHandler {
	q := db.New(pg)

	return &UserHandler{
		ctx,
		l,
		q,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser db.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if dbErr := uh.q.CreateUser(uh.ctx, newUser); dbErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(dbErr)
		return
	}

	jsonInBytes, jsonErr := json.Marshal(map[string]any{
		"success": true,
		"message": "Successfully created User",
	})
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonInBytes))
}

func (uh *UserHandler) GetManyUser(w http.ResponseWriter, r *http.Request) {
	users, dbErr := uh.q.GetUsers(uh.ctx)
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		uh.l.Println(dbErr)
		return
	}

	jsonInBytes, jsonErr := json.Marshal(users)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonInBytes))
}

func (uh *UserHandler) GetOneUserById(w http.ResponseWriter, r *http.Request) {
	userId, paramErr := uuid.Parse(chi.URLParam(r, "userId"))
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(paramErr)
		return
	}

	user, dbErr := uh.q.GetOneUserById(uh.ctx, userId)
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		uh.l.Println(dbErr)
		return
	}

	render.JSON(w, r, utils.CreateBaseResponse(
		true,
		"User found",
		user,
	))

	return
}
