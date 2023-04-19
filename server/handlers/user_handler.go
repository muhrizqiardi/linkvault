package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"server/db"
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
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(jsonInBytes))
}
