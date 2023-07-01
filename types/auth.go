package types

type LoginDTO struct {
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"password"`
}

type SignupDTO struct {
    LoginDTO
    Name string `json:"name" validate:"required,min=3"`
}

type UserRespone struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"-"`
}
