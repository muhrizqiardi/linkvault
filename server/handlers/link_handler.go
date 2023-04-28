package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/dtos"
	"server/services"
	"server/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LinkHandler struct {
	ctx         context.Context
	l           *log.Logger
	linkService services.LinkService
}

func NewLinkHandler(ctx context.Context, l *log.Logger, pg *sqlx.DB) *LinkHandler {
	linkService := services.NewLinkService(ctx, l, pg)
	return &LinkHandler{
		ctx,
		l,
		*linkService,
	}
}

// Create link on default folder
//
// @Summary Create link
// @Tags link
// @Accept json
// @Procedure json
// @Param userId path uuid true "User/Owner ID"
// @Success 201 {object} utils.BaseResponse[entities.LinkEntity]
// @Failure 4xx {object} utils.BaseResponse[any]
// @Failure 5xx {object} utils.BaseResponse[any]
// @Router /users/{userId}/links
func (l *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	userId, parseUserIdErr := uuid.Parse(chi.URLParam(r, "userId"))
	if parseUserIdErr != nil {
		utils.BaseResponseWriter[any](
			w,
			http.StatusBadRequest,
			false,
			"Bad Request",
			nil,
		)
		return
	}
	folderId, parseFolderIdErr := uuid.Parse(chi.URLParam(r, "folderId"))
	if parseFolderIdErr != nil {
		utils.BaseResponseWriter[any](
			w,
			http.StatusBadRequest,
			false,
			"Bad Request",
			nil,
		)
		return
	}

	var payload dtos.CreateLinkDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BaseResponseWriter[any](
			w,
			http.StatusBadRequest,
			false,
			"Bad Request",
			nil,
		)
		return
	}

	newLink, createErr := l.linkService.Create(userId, folderId, payload)
	if createErr != nil {
		utils.BaseResponseWriter[any](
			w,
			http.StatusBadRequest,
			false,
			"Bad Request",
			nil,
		)
		return
	}

	utils.BaseResponseWriter(
		w,
		http.StatusCreated,
		true,
		"Successfully created link",
		newLink,
	)
	return
}
