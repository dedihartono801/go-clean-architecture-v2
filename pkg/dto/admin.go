package dto

type AdminCreateDto struct {
	Name     string `json:"name" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}
