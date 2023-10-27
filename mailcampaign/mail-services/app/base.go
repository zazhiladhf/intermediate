package app

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// repo := NewPosrgresRepository(db)
	svc := NewService()
	handler := NewHandler(svc)

	// panggil fungsi send mail
	// err := SendMail(to, cc, subject, message)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("success send mail to", append(to, cc...))

	router.POST("/send", handler.SendMail)

}
