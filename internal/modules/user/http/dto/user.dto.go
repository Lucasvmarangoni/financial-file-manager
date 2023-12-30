package dto

type UserInput struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
