package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"server/db"
	"server/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var EXPIRATION_TIME = time.Now().Add(6 * 30 * 24 * time.Hour)

type AuthHandler struct {
	ctx context.Context
	l   *log.Logger
	q   *db.Queries
}

func NewAuthHandler(ctx context.Context, l *log.Logger, pg *sql.DB) *AuthHandler {
	q := db.New(pg)

	return &AuthHandler{
		ctx: ctx,
		l:   l,
		q:   q,
	}
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

//	@Summary	Log in to user account
//	@Tags		auth
//	@Produce	json
//	@Param		data	body		LoginParams	true	"Login params"
//	@Success	200		{object}	utils.BaseResponse[string]
//	@Router		/auth [post]
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var param LoginParams
	if decErr := json.NewDecoder(r.Body).Decode(&param); decErr != nil {
		http.Error(w, decErr.Error(), http.StatusBadRequest)
		ah.l.Println(decErr.Error())
		return
	}

	user, dbErr := ah.q.GetOneUserByEmail(ah.ctx, param.Email)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusNotFound)
		ah.l.Println(dbErr.Error())
		return
	}

	if pwIsCorrect := utils.CheckPasswordIsCorrect(param.Password, user.Password); !pwIsCorrect {
		http.Error(w, "Password is incorrect", http.StatusUnauthorized)
		ah.l.Println("Incorrect password")
		return
	}

	claims := Claims{
		Email:  param.Email,
		UserId: user.ID.String(),
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
