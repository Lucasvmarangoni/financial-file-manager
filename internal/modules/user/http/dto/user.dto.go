package dto

type UserInput struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

type AuthenticationInput struct {
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}