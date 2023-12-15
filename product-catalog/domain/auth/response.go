package auth

type Response struct {
	HttpCode  int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Error     string      `json:"error,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
}

type Payload struct {
	AccessToken string `json:"access_token,omitempty"`
	Role        string `json:"role,omitempty"`
}

// func ApiResponse(c *fiber.Ctx, httpCode int, success bool, message string, err string, errorCode string) (resp Response) {
// 	c = c.Status(httpCode)

// 	jsonResponse := Response{
// 		HttpCode:  httpCode,
// 		Success:   success,
// 		Message:   message,
// 		Error:     err,
// 		ErrorCode: errorCode,
// 	}
// 	return jsonResponse
// }

// func ResponseError(c *fiber.Ctx, err error) error {
// 	switch err {
// 	case ErrEmailEmpty:
// 		return ApiResponse(c, "bad request", err.Error(), http.StatusBadRequest, ErrCodeEmailEmpty, nil)
// 	case ErrInvalidEmail:
// 		return ApiResponse(c, "bad request", err.Error(), http.StatusBadRequest, ErrCodeInvalidEmail, nil)
// 	case ErrPasswordEmpty:
// 		return ApiResponse(c, "bad request", err.Error(), http.StatusBadRequest, ErrCodePasswordEmpty, nil)
// 	case ErrInvalidPassword:
// 		return ApiResponse(c, "bad request", err.Error(), http.StatusBadRequest, ErrCodeInvalidPassword, nil)
// 	case ErrDuplicateEmail:
// 		// log.Println("err:", err)
// 		return ApiResponse(c, "duplicate entry", err.Error(), http.StatusConflict, ErrCodeDuplicateEmail, nil)
// 	case ErrInternalServer:
// 		return ApiResponse(c, "internal server error", "unknown error", http.StatusInternalServerError, ErrCodeInternalServer, nil)
// 	default:
// 		log.Println("err:", err)
// 		return ApiResponse(c, "internal server error", "error repository", http.StatusInternalServerError, ErrCodeInternalServer, nil)
// 	}
// }

// func ResponseSuccess(c *fiber.Ctx, success bool, message string, statusCode int, payload interface{}) error {
// 	resp := Response{
// 		Success:   success,
// 		Message:   message,
// 		Error:     "",
// 		ErrorCode: "",
// 		// Payload:   Payload{},
// 	}
// 	c = c.Status(statusCode)
// 	return c.JSON(resp)
// }
