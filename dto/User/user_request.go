package usersdto

type UserRequest struct {
	Id       int
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
