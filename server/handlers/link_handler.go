package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/dtos"
	"server/services"
	"server/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type LinkHandler struct {
	ctx         context.Context
	l           *log.Logger
	linkService services.LinkService
}

func NewLinkHandler(ctx context.Context, l *log.Logger, linkService services.LinkService) *LinkHandler {
	return &LinkHandler{
		ctx,
		l,
		linkService,
	}
}

// Create link on default folder
//
//	@Summary	Create link
//	@Tags		link
//	@Accept		json
//	@Procedure	json
//	@Param		folderId	path		string									true	"Folder ID"
//	@Param		payload		body		dtos.CreateLinkDto						true	"Payload"
//	@Success	201			{object}	utils.BaseResponse[entities.LinkEntity]	"Successfully created a link"
//	@Failure	400			{object}	utils.BaseResponse[any]					"Bad Request"
//	@Failure	500			{object}	utils.BaseResponse[any]					"Internal Server Error"
//	@Security	Bearer
//	@Router		/folders/{folderId}/links [post]
func (l *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	userClaim, ok := r.Context().Value("user").(*Claims)
	if !ok {
		utils.BaseResponseWriter[any](w, http.StatusUnauthorized, false, "Unauthorized", nil)
		l.l.Println("Invalid JWT")
		return
	}
	folderId, parseFolderIdErr := uuid.Parse(chi.URLParam(r, "folderId"))
	if parseFolderIdErr != nil {
		utils.BaseResponseWriter[any](w, http.StatusBadRequest, false, "Bad Request", nil)
		l.l.Println(parseFolderIdErr.Error())
		return
	}

	validate := validator.New()
	var payload dtos.CreateLinkDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BaseResponseWriter[any](w, http.StatusBadRequest, false, "Bad Request", nil)
		l.l.Println(err)
		return
	}
	if err := validate.Struct(payload); err != nil {
		utils.BaseResponseWriter[any](w, http.StatusBadRequest, false, "Bad Request", nil)
		l.l.Println(err.Error())
		return
	}

	newLink, createErr := l.linkService.Create(uuid.MustParse(userClaim.UserId), folderId, payload)
	if createErr != nil {
		utils.BaseResponseWriter[any](w, http.StatusInternalServerError, false, "Internal Server Error", nil)
		l.l.Println(createErr)
		return
	}

	utils.BaseResponseWriter(w, http.StatusCreated, true, "Successfully created link", newLink)
	return
}

// Get many link belongs to a user
//
//	@Summary	Get many link
//	@Tags		link
//	@Accept		json
//	@Procedure	json
//	@Param		title	query		string										false	"Search matching title"
//	@Param		excerpt	query		string										false	"Search matching excerpt"
//	@Param		orderBy	query		string										false	"Order by title, created date, or modified date"	Enum(title_ASC, title_DESC, createdAt_ASC, createdAt_DESC, updatedAt_ASC, updatedAt_DESC)	default(updatedAt_DESC)
//	@Param		limit	query		int											false	"Limit every page"									default(10)
//	@Param		page	query		int											false	"Page count"										default(1)
//	@Success	200		{object}	[]utils.BaseResponse[entities.LinkEntity]	"Successfully created user"
//	@Failure	400		{object}	utils.BaseResponse[any]						"Bad Request"
//	@Failure	500		{object}	utils.BaseResponse[any]						"Internal Server Error"
//	@Security	Bearer
//	@Router		/links [get]
func (l *LinkHandler) GetManyLinks(w http.ResponseWriter, r *http.Request) {
	userClaim, _ := r.Context().Value("user").(*Claims)
	title := r.URL.Query().Get("title")
	excerpt := r.URL.Query().Get("excerpt")
	orderBy := r.URL.Query().Get("orderBy")
	if err := utils.ValidateEnumString(orderBy, "title_ASC", "title_DESC", "createdAt_ASC", "createdAt_DESC", "updatedAt_ASC", "updatedAt_DESC"); err != nil {
		utils.BaseResponseWriter[any](w, http.StatusBadRequest, false, "Bad Request", nil)
		l.l.Println(err.Error())
		return
	}
	defaultLimit := 10
	defaultPage := 1
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = defaultLimit
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = defaultPage
	}
	links, err := l.linkService.GetManyBelongsToUser(uuid.MustParse(userClaim.UserId), title, excerpt, orderBy, limit, page)
	if err != nil {
		utils.BaseResponseWriter[any](w, http.StatusInternalServerError, false, "Internal Server Error", nil)
		l.l.Println(err.Error())
		return
	}
	if len(links) == 0 {
		utils.BaseResponseWriter[any](w, http.StatusNotFound, false, "Link(s) not found", nil)
		l.l.Println("Link(s) not found")
		return
	}

	utils.BaseResponseWriter(w, http.StatusOK, true, "Link(s) found", links)
	return
}
