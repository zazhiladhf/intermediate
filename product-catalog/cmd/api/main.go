package main

import (
	"log"
	"product-catalog/config"
	"product-catalog/domain/auth"
	"product-catalog/domain/category"
	"product-catalog/domain/product"
	"product-catalog/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	router := fiber.New(fiber.Config{
		AppName: config.Cfg.App.Name,
		// Prefork: true,
	})

	dbSqlx, err := database.ConnectPostgresSqlx(config.Cfg.DB)
	if err != nil {
		log.Println("error when to try migration db with error :", err.Error())
		// panic(err)
	}

	log.Println("running db migration")
	err = database.Migrate(dbSqlx)
	if err != nil {
		log.Println("migration failed with error:", err)
		panic(err)
	}
	log.Println("migration done")

	auth.Run(router, dbSqlx)
	category.RegisterCategoriesRouter(router, dbSqlx)
	product.RegisterServiceProduct(router, dbSqlx)

	// redis(5 * time.Second)

	router.Listen(config.Cfg.App.Port)
}

// func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if !strings.Contains(authHeader, "Bearer") {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		tokenString := ""
// 		arrayToken := strings.Split(authHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		token, err := authService.ValidateToken(tokenString)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		claim, ok := token.Claims.(jwt.MapClaims)

// 		if !ok || !token.Valid {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		userID := int(claim["user_id"].(float64))

// 		user, err := userService.GetUserByID(userID)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		c.Set("currentUser", user)
// 	}

// }

// func redis(timeout time.Duration) {
// 	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
// 	defer cancel()

// 	client, err := database.ConnectRedis(config.Cfg.Redis)
// 	if err != nil {
// 		log.Println("error when to try migration db with error :", err.Error())
// 		// panic(err)
// 	}

// 	// client, err := database.ConnectRedis(ctx, "localhost:6377", "")
// 	// if err != nil {
// 	// 	log.Println("db not connected with error", err.Error())
// 	// 	return
// 	// }

// 	if client == nil {
// 		log.Println("db not connected with unknown error")
// 		return
// 	}

// 	err = client.Set(ctx, "token", "ini-token-user", timeout).Err()
// 	if err != nil {
// 		log.Println("error when try to set data to redis with message :", err.Error())
// 		return
// 	}

// 	cmd := client.Get(ctx, "token-user02")

// 	res, err := cmd.Result()
// 	if err != nil {
// 		log.Println("error when try to get data from redis with message :", err.Error())
// 		return
// 	}

// 	client.Del(ctx, "token-user02")
// 	ttl, err := client.TTL(ctx, "token-user01").Result()
// 	if err != nil {
// 		log.Println("error when try to get data from redis with message :", err.Error())
// 		return
// 	}

// 	log.Println("ttl token-user01", ttl)
// 	log.Println("isi token token-user02 adalah", res)

// 	log.Println("redis connected")
// }
