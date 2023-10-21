package main

import (
	"goframework/config"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Request struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

var data = []Request{
	{
		Id:      1,
		Name:    "string",
		Email:   "string",
		Address: "string",
	},
	{
		Id:      2,
		Name:    "zazhil",
		Email:   "string",
		Address: "string",
	},
}

func main() {

	router := gin.New()

	router.Use(Trace())

	router.GET("/users", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"status_code": 200,
			"message":     "get all success",
			"payload":     data,
		})
	})

	router.POST("/users", func(ctx *gin.Context) {
		var req = Request{}
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		req.Id = len(data) + 1

		data = append(data, req)

		// auth := ctx.GetHeader("Authorization")

		ctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"status_code": 201,
			"message":     "created success",
		})
	})

	router.PUT("/users/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")
		for _, v := range data {
			idParam, _ := strconv.Atoi(id)
			if v.Id == idParam {
				var req = Request{}
				err := ctx.ShouldBind(&req)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}
				req.Id = idParam
				data[idParam-1] = req

				ctx.JSON(http.StatusOK, gin.H{
					"success":     true,
					"status_code": 200,
					"message":     "update all success",
				})
			}
		}
	})

	router.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for _, v := range data {
			idParam, _ := strconv.Atoi(id)
			if v.Id == idParam {
				idParam = idParam - 1
				data = append(data[:idParam], data[idParam+1:]...)

				ctx.JSON(http.StatusOK, gin.H{
					"success":     true,
					"status_code": 200,
					"message":     "delete success",
				})
			}
		}
	})

	router.Run(config.APP_PORT)
}

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := uuid.NewString()
		ctx.Header("X-Trace-ID", traceId)
		// ctx.Set()

		message := "incoming request"
		method := "GET"
		uri := "/users"
		log.Printf("message=%s method=%s uri=%s trace_id=%s", message, method, uri, traceId)

		ctx.Next()

		message = "error when try to get users with no data"
		log.Printf("message=%s method=%s uri=%s trace_id=%s", message, method, uri, traceId)

		ctx.Next()

		message = "finish request"
		log.Printf("message=%s method=%s uri=%s trace_id=%s", message, method, uri, traceId)

		// log.Println("Middleware 1: After Request")
	}
}
