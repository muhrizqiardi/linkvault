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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if encodeErr := json.NewEncoder(w).Encode(
		utils.CreateBaseResponse(
			true,
			"User created",
			nil,
		),
	); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		uh.l.Println(encodeErr)
		return
	}
}

func (uh *UserHandler) GetManyUser(w http.ResponseWriter, r *http.Request) {
	users, dbErr := uh.q.GetUsers(uh.ctx)
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		uh.l.Println(dbErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if encodeErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(
		true,
		"User(s) found",
		users,
	)); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		uh.l.Println(encodeErr)
		return
	}

	return
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(
		true,
		"User found",
		user,
	)); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		uh.l.Println(encodeErr)
		return
	}

	return
}

func (uh *UserHandler) UpdateOneUserById(w http.ResponseWriter, r *http.Request) {
	userId, paramErr := uuid.Parse(chi.URLParam(r, "userId"))
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(paramErr)
		return
	}

	var payload db.UpdateOneUserByIdParams
	if decErr := json.NewDecoder(r.Body).Decode(&payload); decErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(decErr)
		return
	}

	if dbErr := uh.q.UpdateOneUserById(uh.ctx, db.UpdateOneUserByIdParams{
		ID:       userId,
		Email:    payload.Email,
		Password: payload.Password,
		FullName: payload.FullName,
	}); dbErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(dbErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(true, "User updated", nil)); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		uh.l.Println(encodeErr)
		return
	}

	return
}

func (uh *UserHandler) DeleteOneUserById(w http.ResponseWriter, r *http.Request) {
	userId, paramErr := uuid.Parse(chi.URLParam(r, "userId"))
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(paramErr)
		return
	}

	if dbErr := uh.q.DeleteOneUserById(uh.ctx, userId); dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		uh.l.Println(dbErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(true, "User deleted", nil)); encErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		uh.l.Println(encErr)
		return
	}

	return
}
