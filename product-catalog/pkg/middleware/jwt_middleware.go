package middleware

import (
	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:   []byte("secret"),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}

// func AuthMiddleware() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		var err error
// 		var resp = auth.Auth{}

// 		authorization := c.Get("Authorization")

// 		if !strings.Contains(authorization, "Bearer") {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": true,
// 				"msg":   err.Error(),
// 			})
// 		}

// 		tokenString := ""
// 		arrayToken := strings.Split(authorization, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		claims, err := jwt.ValidateToken(tokenString)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": true,
// 				"msg":   err.Error(),
// 			})
// 		}

// 		// if !ok || !token.Valid {
// 		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		// 		"error": true,
// 		// 		"msg":   err.Error(),
// 		// 	})
// 		// }

// 		resp.Email = claims["email"].(string)

// 		auth, err := auth.AuthService.GetAuthByEmail(auth.AuthService, c.UserContext(), resp.Email)
// 		if err != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": true,
// 				"msg":   err.Error(),
// 			})
// 		}

// 		c.Set("currentUser", auth.Email)

// 		return c.Status(http.StatusOK).JSON(fiber.Map{
// 			"user": resp,
// 		})
// 	}
// }
