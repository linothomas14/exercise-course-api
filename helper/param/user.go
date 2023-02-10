package param

type UserUpdate struct {
	Name     string `json:"name"`
	Email    string `json:"email" `
	Password string `json:"password" validate:"required,min=6"`
}
