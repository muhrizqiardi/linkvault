package dtos

type CreateUserDto struct {
	Email    string `json:"email" validation:"required,email" format:"email"`
	FullName string `json:"full_name" validation:"required"`
	Password string `json:"password" validation:"required,min=8" minLength:"8"`
}

type UpdateUserDto struct {
	Email    string `json:"email" validation:"required,email" format:"email"`
	Password string `json:"password" validation:"required,min=8" minLength:"8"`
	FullName string `json:"full_name"`
}
