package app

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) HandlerRequest(ctx *gin.Context) {
	url := "http://localhost:4001"
	var reqBody RequestBody
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	reqData := RequestBody{
		From:    reqBody.From,
		To:      reqBody.To,
		Subject: reqBody.Subject,
		Message: reqBody.Message,
		Type:    reqBody.Type,
	}

	// ubah request data menjadi sebuah []byte
	byteReq, err := json.Marshal(reqData)
	if err != nil {
		panic(err)
	}

	// data yg dikirim via http, adalah sebuah bytes buffer
	// jadi perlu kita ubah dari []byte ke bytes buffer
	buf := bytes.NewBuffer(byteReq)

	resp, err := http.Post(url+"/send", "application/json", buf)
	if err != nil {
		panic(err)
	}

	// close response body jika sudah di baca
	defer resp.Body.Close()

	// proses validasi status code
	// karena dari service1 nge return 200, maka perlu di pastikan bahwa
	// yang di porses hanya yg status code nya 200
	if resp.StatusCode != http.StatusOK {
		log.Println("ada error gak?")
		responseBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error saat parsing response body with error : %v", err.Error())
			return
		}
		log.Println("Response :", string(responseBytes))
		return
	}

	// proses baca response body.
	// jadi hasilnya akan berupa sebuah []byte, yang mana bisa
	// kita jadikan ke sebuah map atau struct nantinya
	// dengan cara melakukan json.Unmarshal()
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error saat parsing response body with error : %v", err.Error())
		return
	}

	log.Println("Response :", string(responseBytes))

	// Handle response
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Request sent successfully",
	})

}
