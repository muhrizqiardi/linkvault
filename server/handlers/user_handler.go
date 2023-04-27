package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/dtos"
	"server/entities"
	"server/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	ctx context.Context
	l   *log.Logger
	pg  *sqlx.DB
}

func NewUserHandler(ctx context.Context, l *log.Logger, pg *sqlx.DB) *UserHandler {
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
	if queryErr := uh.pg.QueryRowx(createUserQuery, &newUser.Email, &newUser.FullName, &hashedPassword).StructScan(&userResult); queryErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Failed to create user", nil))
		uh.l.Println(queryErr.Error())
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
	dbErr := uh.pg.SelectContext(uh.ctx, &users, getManyUsersQuery)
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "User(s) not found", nil))
		uh.l.Println(dbErr.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if encodeErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(
		true,
		"User(s) found",
		users,
	)); encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(dbErr.Error())
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
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Bad Request", nil))
		uh.l.Println(paramErr.Error())
		return
	}

	userClaim, ok := r.Context().Value("user").(*Claims)
	if !ok || userClaim.UserId != userId.String() {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "User not found", nil))
		uh.l.Println("Unauthorized access")
		return
	}

	getOneUserByIdQuery := `
		select * from public.users 
		    where id = $1;
	`

	var user entities.UserEntity
	if dbErr := uh.pg.QueryRowx(getOneUserByIdQuery, userId).StructScan(&user); dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "User not found", nil))
		uh.l.Println("User not found")
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
//	@Param		userId	path	string				true	"User id"
//	@Param		data	body	dtos.CreateUserDto	true	"create user param"
//	@Produce	json
//	@Success	200	{object}	utils.BaseResponse[entities.UserEntity]
//	@Router		/users/{userId} [patch]
//	@Security	Bearer
func (uh *UserHandler) UpdateOneUserById(w http.ResponseWriter, r *http.Request) {
	userId, paramErr := uuid.Parse(chi.URLParam(r, "userId"))
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Bad Request", nil))
		uh.l.Println(paramErr.Error())
		return
	}

	userClaim, ok := r.Context().Value("user").(*Claims)
	if !ok && userClaim.UserId != userId.String() {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Unauthorized", nil))
		uh.l.Println("Unauthorized")
		return
	}

	var payload dtos.UpdateUserDto
	if decErr := json.NewDecoder(r.Body).Decode(&payload); decErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(decErr.Error())
		return
	}

	hashedPassword, pwErr := utils.HashPassword(payload.Password)
	if pwErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(pwErr.Error())
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
	if dbErr := uh.pg.QueryRowx(updateOneUserByIdQuery, userId, payload.Email, hashedPassword, payload.FullName).StructScan(&updatedUser); dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(dbErr.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if encErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(true, "User updated", updatedUser)); encErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(encErr.Error())
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
	w.Header().Set("Content-Type", "application/json")

	userId, paramErr := uuid.Parse(chi.URLParam(r, "userId"))
	if paramErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Bad Request", nil))
		uh.l.Println(paramErr.Error())
		return
	}

	userClaim, ok := r.Context().Value("user").(*Claims)
	if !ok && userClaim.UserId != userId.String() {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Unauthorized", nil))
		uh.l.Println(paramErr.Error())
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
	if dbErr := uh.pg.QueryRowx(deleteOneUserByIdQuery, userId).StructScan(&deletedUser); dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(dbErr.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if encErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(true, "User deleted", deletedUser)); encErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		uh.l.Println(encErr)
		return
	}

	return
}
