package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/dtos"
	_ "server/entities"
	"server/services"
	"server/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type FolderHandler struct {
	ctx           context.Context
	l             *log.Logger
	folderService services.FolderService
}

func NewFolderHandler(ctx context.Context, l *log.Logger, linkService services.FolderService) *FolderHandler {
	return &FolderHandler{
		ctx,
		l,
		linkService,
	}
}

// Create a folder
//
//	@Summary	Create a folder
//	@Tags		folder
//	@Accept		json
//	@Procedure	json
//	@Param		payload	body		dtos.CreateFolderDto						true	"Payload"
//	@Success	201		{object}	utils.BaseResponse[entities.FolderEntity]	"Successfully created a folder"
//	@Failure	400		{object}	utils.BaseResponse[any]						"Bad Request"
//	@Failure	500		{object}	utils.BaseResponse[any]						"Internal Server Error"
//	@Security	Bearer
//	@Router		/folders [post]
func (fh *FolderHandler) CreateFolder(w http.ResponseWriter, r *http.Request) {
	userClaims, ok := r.Context().Value("user").(*Claims)
	if !ok {
		utils.BaseResponseWriter[any](w, http.StatusUnauthorized, false, "Unauthorized", nil)
		fh.l.Println("Invalid JWT")
		return
	}

	validate := validator.New()
	var payload dtos.CreateFolderDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.BaseResponseWriter[any](w, http.StatusBadRequest, false, "Bad Request", nil)
		fh.l.Println(err.Error())
		return
	}
	if err := validate.Struct(payload); err != nil {
		utils.BaseResponseWriter[any](w, http.StatusBadRequest, false, "Bad Request", nil)
		fh.l.Println(err.Error())
		return
	}

	newFolder, err := fh.folderService.Create(uuid.MustParse(userClaims.UserId), payload)
	if err != nil {
		utils.BaseResponseWriter[any](w, http.StatusInternalServerError, false, "Internal Server Error", nil)
		fh.l.Println(err.Error())
		return
	}

	utils.BaseResponseWriter[any](w, http.StatusCreated, true, "Successfully created a folder", newFolder)
	return
}

// Get many folders belongs to user
//
//	@Summary	Get many folders belongs to user
//	@Tags		folder
//	@Accept		json
//	@Procedure	json
//	@Param		orderBy	query		string										false	"Order by title, created date, or modified date"	Enum(title_ASC, title_DESC, createdAt_ASC, createdAt_DESC, updatedAt_ASC, updatedAt_DESC)	default(updatedAt_DESC)
//	@Param		limit	query		int											false	"Limit every page"									default(10)
//	@Param		page	query		int											false	"Page count"										default(1)
//	@Success	201		{object}	utils.BaseResponse[entities.FolderEntity]	"Folder(s) found"
//	@Failure	400		{object}	utils.BaseResponse[any]						"Bad Request"
//	@Failure	500		{object}	utils.BaseResponse[any]						"Internal Server Error"
//	@Security	Bearer
//	@Router		/folders [get]
func (fh *FolderHandler) GetManyFoldersBelongsToUser(w http.ResponseWriter, r *http.Request) {
	userClaims, _ := r.Context().Value("user").(*Claims)
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
	orderBy := r.URL.Query().Get("orderBy")
	if err := utils.ValidateEnumString(orderBy, "title_ASC", "title_DESC", "createdAt_ASC", "createdAt_DESC", "updatedAt_ASC", "updatedAt_DESC"); err != nil {
		utils.BaseResponseWriter[any](w, http.StatusBadRequest, false, "Bad Request", nil)
		fh.l.Println(err.Error())
		return
	}

	folders, err := fh.folderService.GetManyBelongsToUser(uuid.MustParse(userClaims.UserId), orderBy, limit, page)
	if err != nil {
		utils.BaseResponseWriter[any](w, http.StatusBadRequest, false, "Bad Request", nil)
		fh.l.Println(err.Error())
		return
	}

	utils.BaseResponseWriter(w, http.StatusOK, true, "Folder(s) found", folders)
	return
}