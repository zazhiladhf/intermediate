package auth

type registerRequest struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

// type registerResponse struct {
// 	Id       int    `json:"id"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	Role     string `json:"role"`
// }

// func newRegisterResponse(item Auth) registerResponse {
// 	return registerResponse{
// 		Id:       item.Id,
// 		Email:    item.Email,
// 		Password: item.Password,
// 		Role:     item.Role,
// 	}
// }

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type loginResponse struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
