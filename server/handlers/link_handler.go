package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
)

type LinkHandler struct {
	ctx context.Context
	l   *log.Logger
	pg  *sql.DB
}

func NewLinkHandler(ctx context.Context, l *log.Logger, pg *sql.DB) *LinkHandler {
	return &LinkHandler{
		ctx,
		l,
		pg,
	}
}

func (l *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	// createLinkQuery := `
	// 	insert into public.links (url, excerpt, cover_url, owner_id, folder_id)
	//    		values ($1, $2, $3, $4, $5)
	//    		returning id, url, excerpt, cover_url, owner_id, folder_id;
	// `
	// w.Header().Set("Content-Type", "application/json")

	// _ := chi.URLParam(r, "userId")
	// var newLinkParams dtos.CreateLinkDto
	// if decErr := json.NewDecoder(r.Body).Decode(&newLinkParams); decErr != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Bad request", nil))
	// 	return
	// }

	// var queryResult dtos.CreateLinkDto
	// if err := l.pg.QueryRowContext(l.ctx, createLinkQuery, newLinkParams.Url, newLinkParams.Excerpt, newLinkParams).Scan(&queryResult); err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Failed to create link", nil))
	// 	return
	// }

}
