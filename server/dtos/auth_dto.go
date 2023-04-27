package dtos

type AuthLoginDto struct {
	Email    string `json:"email" validation:"required,email" format:"email"`
	Password string `json:"password" validation:"required,min=8" minLength:"8"`
}
