package dto

type UserCreateDto struct {
	Name  string `json:"name" `
	Email string `json:"email"`
}

type UserUpdateDto struct {
	Name  string `json:"name" `
	Email string `json:"email"`
}
