package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{
		svc: svc,
	}
}

func (h Handler) SendMail(ctx *gin.Context) {
	var cc []string
	var reqBody RequestBody

	err := ctx.ShouldBind(&reqBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// setup email tujuan
	// to := []string{"user1@gmail.com", "user2@gmail.com"}
	// setup cc
	// cc := []string{"user3@gmail.com"}

	// subject := "Test Mail"
	// messagenya menjadi HTML
	// message := `
	// <html>
	// 	<body>
	// 		<h1> Hello From NooBeeID</h1>
	// 		<button class="btn btn-primary ">Click Me</button>
	// 	</body>
	// </html>
	// `

	// panggil fungsi send mail
	// err = sendMailGoMail(to, cc, subject, message)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("success send mail to", append(to, cc...))

	reqData := RequestBody{
		From:    reqBody.From,
		To:      reqBody.To,
		Subject: reqBody.Subject,
		Message: reqBody.Message,
		Type:    reqBody.Type,
	}

	err = h.svc.SendMailService(ctx, reqData.To, cc, reqData.Subject, reqData.Message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": reqData.Message,
		// "error":   "string", // if no error, this attribute will be remove
	})
}
