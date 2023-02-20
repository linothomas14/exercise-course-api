package param

type AdminUpdate struct {
	ID       uint32
	Name     string `json:"name"`
	Email    string `json:"email" `
	Password string `json:"password"`
}
