package auth

type register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newRegisterResponse(item Auth) registerResponse {
	return registerResponse{
		Id:       item.Id.String(),
		Email:    item.Email,
		Password: item.Password,
	}
}

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
