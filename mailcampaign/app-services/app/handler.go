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
	// svc Service
}

func NewHandler() Handler {
	return Handler{
		// svc: svc,
	}
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

	// ubah request data menjadi sebuah []byte
	byteReq, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteReq))
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// req.Header.Set("Content-Type", "application/json")

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
	if resp.StatusCode != http.StatusCreated {
		log.Println("ada error nih")
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
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Request sent successfully",
		// "error":   "string", // if no error, this attribute will be remove
	})

}
