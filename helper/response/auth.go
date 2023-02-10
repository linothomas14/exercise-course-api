package response

type RegisterResponse struct {
	ID    uint32 `json:"id" `
	Email string `json:"email" `
	Name  string `json:"name" `
}
