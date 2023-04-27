package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"server/dtos"
	"server/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

var EXPIRATION_TIME = time.Now().Add(6 * 30 * 24 * time.Hour)

type AuthHandler struct {
	ctx context.Context
	l   *log.Logger
	pg  *sqlx.DB
}

func NewAuthHandler(ctx context.Context, l *log.Logger, pg *sqlx.DB) *AuthHandler {

	return &AuthHandler{
		ctx: ctx,
		l:   l,
		pg:  pg,
	}
}

type Claims struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

// @Summary	Log in to user account
// @Tags		auth
// @Produce	json
// @Param		data	body		dtos.AuthLoginDto	true	"Login params"
// @Success	200		{object}	utils.BaseResponse[string]
// @Router		/auth [post]
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var param dtos.AuthLoginDto
	if decErr := json.NewDecoder(r.Body).Decode(&param); decErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Bad Request", nil))
		ah.l.Println(decErr.Error())
		return
	}

	getOneUserByEmailQuery := `
		select id, email, password 
			from public.users
			where
				email = $1;
	`
	var foundUser struct {
		Id       string
		Email    string
		Password string
	}
	if scanErr := ah.pg.QueryRow(getOneUserByEmailQuery, param.Email).Scan(&foundUser.Id, &foundUser.Email, &foundUser.Password); scanErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Internal Server Error", nil))
		ah.l.Println(scanErr.Error())
		return
	}

	if pwIsCorrect := utils.CheckPasswordIsCorrect(param.Password, foundUser.Password); !pwIsCorrect {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(utils.CreateBaseResponse[any](false, "Unauthorized", nil))
		ah.l.Println("Incorrect password")
		return
	}

	claims := Claims{
		Email:  param.Email,
		UserId: foundUser.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(EXPIRATION_TIME),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, signErr := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if signErr != nil {
		http.Error(w, signErr.Error(), http.StatusInternalServerError)
		ah.l.Println(signErr.Error())
		return
	}

	//http.SetCookie(w, &http.Cookie{
	//	Name: "token",
	//	Value: tokenString,
	//	Expires: EXPIRATION_TIME,
	//})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if encErr := json.NewEncoder(w).Encode(utils.CreateBaseResponse(true, "Log in success", tokenString)); encErr != nil {
		http.Error(w, signErr.Error(), http.StatusInternalServerError)
		ah.l.Println(encErr.Error())
		return
	}

	return
}
