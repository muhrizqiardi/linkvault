package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"server/dtos"
	"server/entities"
	"server/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	ctx context.Context
	l   *log.Logger
	pg  *sql.DB
}

func NewUserHandler(ctx context.Context, l *log.Logger, pg *sql.DB) *UserHandler {
	return &UserHandler{
		ctx,
		l,
		pg,
	}
}

//	@Router		/users [post]
//	@Tags		users
//	@Summary	Create user
//	@Produce	json
//	@Param		data	body		dtos.CreateUserDto	true	"create user param"
//	@Success	201		{object}	utils.BaseResponse[entities.UserEntity]
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser dtos.CreateUserDto
	if decErr := json.NewDecoder(r.Body).Decode(&newUser); decErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(decErr.Error())
		return
	}

	hashedPassword, pwErr := utils.HashPassword(newUser.Password)
	if pwErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(pwErr.Error())
	}

	createUserQuery := `
		insert into public.users (email, full_name, password) 
			values ($1, $2, $3)
			returning id, email, full_name, password, created_at, updated_at;`
	var userResult entities.UserEntity
	if qErr := uh.pg.QueryRow(
		createUserQuery,
		newUser.Email,
		newUser.FullName,
		hashedPassword,
	).Scan(&userResult.Id, &userResult.Email, &userResult.FullName, &userResult.Password, &userResult.CreatedAt, &userResult.UpdatedAt); qErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Failed to create user", nil))
		uh.l.Println(qErr.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if encErr := json.NewEncoder(w).Encode(
		utils.CreateBaseResponse(
			true,
			"User created",
			userResult,
		),
	); encErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "User not found", nil))
		uh.l.Println(encErr.Error())
		return
	}
}

//	@Router		/users [get]
//	@Tags		users
//	@Summary	Get many users
//	@Produce	json
//	@Success	200	{object}	utils.BaseResponse[[]entities.UserEntity]
func (uh *UserHandler) GetManyUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getManyUsersQuery := `
		select id, email, full_name, password, created_at, updated_at from public.users;
	`
	var users []entities.UserEntity
	queryRows, dbErr := uh.pg.QueryContext(uh.ctx, getManyUsersQuery)
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "User not found", nil))
		return
	}
	queryRows.Scan(&users)

	w.WriteHeader(http.StatusCreated)
	if encodeErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(
		true,
		"User(s) found",
		users,
	)); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		return
	}

	return
}

//	@Summary	Get one user by ID
//	@Tags		users
//	@Param		userId	path	string	true	"User id"
//	@Produce	json
//	@Success	200	{object}	utils.BaseResponse[entities.UserEntity]
//	@Router		/users/{userId} [get]
//	@Security	Bearer
func (uh *UserHandler) GetOneUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, paramErr := uuid.Parse(chi.URLParam(r, "userId"))
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(paramErr)
		return
	}

	userClaim, ok := r.Context().Value("user").(*Claims)
	if !ok && userClaim.UserId != userId.String() {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println("Unauthorized")
		return
	}

	getOneUserByIdQuery := `
		select * from public.users 
		    where id = $1;
	`

	var user entities.UserEntity
	if dbErr := uh.pg.QueryRow(getOneUserByIdQuery, userId).Scan(&user); dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "User not found", nil))
		return
	}

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

//	@Summary	Update one user by ID
//	@Tags		users
//	@Param		userId	path	string	true	"User id"
//	@Produce	json
//	@Success	200	{object}	utils.BaseResponse[entities.UserEntity]
//	@Router		/users/{userId} [put]
//	@Security	Bearer
func (uh *UserHandler) UpdateOneUserById(w http.ResponseWriter, r *http.Request) {
	userId, paramErr := uuid.Parse(chi.URLParam(r, "userId"))
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(paramErr)
		return
	}

	userClaim, ok := r.Context().Value("user").(*Claims)
	if !ok && userClaim.UserId != userId.String() {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println("Unauthorized")
		return
	}

	var payload dtos.UpdateUserDto
	if decErr := json.NewDecoder(r.Body).Decode(&payload); decErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(decErr)
		return
	}

	updateOneUserByIdQuery := `
		update public.users
		    set 
		        email = coalesce($2, email),
		        password = coalesce($3, password),
		        full_name = coalesce($4, full_name),
				updated_at = current_timestamp
		    where 
		        id = $1
			returning
				id,			
				email,
				full_name,
				password,
				created_at,
				updated_at;				
	`

	var updatedUser entities.UserEntity
	if dbErr := uh.pg.QueryRow(updateOneUserByIdQuery, userId).Scan(&updatedUser); dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "User not found", nil))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encodeErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(true, "User updated", updatedUser)); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		uh.l.Println(encodeErr)
		return
	}

	return
}

//	@Summary	Delete one user by ID
//	@Tags		users
//	@Param		userId	path	string	true	"User id"
//	@Produce	json
//	@Success	200	{object}	utils.BaseResponse[entities.UserEntity]
//	@Router		/users/{userId} [delete]
//	@Security	Bearer
func (uh *UserHandler) DeleteOneUserById(w http.ResponseWriter, r *http.Request) {
	userId, paramErr := uuid.Parse(chi.URLParam(r, "userId"))
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println(paramErr)
		return
	}

	userClaim, ok := r.Context().Value("user").(*Claims)
	if !ok && userClaim.UserId != userId.String() {
		w.WriteHeader(http.StatusBadRequest)
		uh.l.Println("Unauthorized")
		return
	}

	deleteOneUserByIdQuery := `
		delete from public.users 
			where 
				id = $1
			returning
				id,			
				email,
				full_name,
				password,
				created_at,
				updated_at;				
	`

	var deletedUser entities.UserEntity
	if dbErr := uh.pg.QueryRow(deleteOneUserByIdQuery, userId).Scan(&deletedUser); dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "User not found", nil))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(true, "User deleted", deletedUser)); encErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		uh.l.Println(encErr)
		return
	}

	return
}
