package main

// import (
// 	"encoding/json"
// 	"log"
// 	"mailcampaign/mail-services/config"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gopkg.in/gomail.v2"
// )

// type RequestBody struct {
// 	From      string `json:"from"`
// 	To 		[]string `json:"to"`
// 	Subject   string `json:"subject"`
// 	Message   string `json:"message"`
// 	Type      string `json:"type"`
// }

// func main() {

// 	config := config.Init()

// 	router := gin.New()

// 	router.GET("/send", postSMTP)

// 	router.Run(config.AppPort)

//     // POST /send
// 	http.HandleFunc("/send", postSMTP)

// 	log.Println("server running at port", config.AppPort)
// 	http.ListenAndServe(config.AppPort, nil)

// 	var input RequestBody

// 	// setup email tujuan
// 	to := []string{"hananayyub@gmail.com", "zazhiladhf@gmail.com"}
// 	// setup cc
// 	cc := []string{"kalajourney1@gmail.com"}

// 	subject := "Test Mail"
// 	message := `
// 	<html>
// 		<body>
// 			<h1> Hello From NooBeeID</h1>
// 			<button class="btn btn-primary ">Click Me</button>
// 		</body>
// 	</html>
// 	`

// 	// panggil fungsi send mail
// 	err := sendMailGoMail(to, cc, subject, message)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println("success send mail to", append(to, cc...))
// }

// // func addUser(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")

// // 	w.WriteHeader(http.StatusCreated)
// // 	json.NewEncoder(w).Encode(users)

// // }

// func sendMailGoMail(to []string, cc []string, subject string, message string) (err error) {

// 	config := config.Init()

// 	// setup gomail message
// 	mailer := gomail.NewMessage()
//     // setting header from
// 	mailer.SetHeader("From", config.SenderName)
//     // setting header to
// 	mailer.SetHeader("To", to...)

//     // setting header CC
// 	for _, ccEmail := range cc {
// 		mailer.SetAddressHeader("Cc", ccEmail, "")
// 	}

//     // setting subject
// 	mailer.SetHeader("Subject", subject)
//     // setting body
//     // kali ini, kita akan menggunakan body HTML agar tampilan dari emailnya lebih menarik
// 	mailer.SetBody("text/html", message)

// 	dialer := gomail.NewDialer(
// 		config.SMTPHost,
// 		config.SMTPPort,
// 		config.AuthEmail,
// 		config.AuthPassword,
// 	)

// 	err = dialer.DialAndSend(mailer)
// 	return
// }

// func postSMTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

//     // makesure method yang digunakan adalah POST
//     if r.Method != http.MethodPost {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"error": "method not allowed",
// 		})
// 	}

// 	var req = RequestBody{}
//     // proses parsing request dari request body
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 	}

//     // generate id
//     req.Id = len(users) + 1

//     // proses menambahkan user
// 	users = append(users, req)

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(users)
// }
