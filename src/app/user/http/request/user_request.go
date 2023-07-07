package request

type UserRequest struct {
	Name    string `json:"name" validate:"required,email"`
	Address string `json:"address" validate:"required"`
	Email   string `json:"email" validate:"required"`
}
