package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

const (
	// config SMTP HOST yaitu gmail
	CONFIG_SMTP_HOST = "smtp.gmail.com"
	// SMTP PORT
	CONFIG_SMTP_PORT = 587
	// Sender Name
	CONFIG_SENDER_NAME = "Zazhil Adhafi <zazhil24@gmail.com>"

	// config untuk authentication nya.
	// gunakan email yg di pakai saat generate app password tadi
	CONFIG_AUTH_EMAIL = "zazhil24@gmail.com"
	// app password yang telah di generate
	CONFIG_AUTH_PASSWORD = "circioqtixpmhowi"
)

func main() {
	// setup email tujuan
	to := []string{"hananayyub@gmail.com", "zazhiladhf@gmail.com"}
	// setup cc
	cc := []string{"kalajourney1@gmail.com"}

	subject := "Test Mail"
	message := `
	<html>
		<body>
			<h1> Hello From NooBeeID</h1>
			<button class="btn btn-primary ">Click Me</button>
		</body>
	</html>
	`

	// panggil fungsi send mail
	// err := sendMail(to, cc, subject, message)
	// if err != nil {
	// 	panic(err)
	// }
    // log.Println("success send mail to",append(to,cc...))

	// panggil fungsi send mail
	err := sendMailGoMail(to, cc, subject, message)
	if err != nil {
		panic(err)
	}
	log.Println("success send mail to", append(to, cc...))
}

// func sendMail(to []string, cc []string, subject, message string) (err error) {
// 	// pada native package smtp, seluruh header ada di dalam body email
// 	// jadi perlu di generate di body email nya
// 	body := "From: " + CONFIG_SENDER_NAME + "\n" +
// 		"To: " + strings.Join(to, ",") + "\n" +
// 		"Cc: " + strings.Join(cc, ",") + "\n" +
// 		"Subject: " + subject + "\n\n" +
// 		message

// 	// Setup untuk authentication
// 	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)

// 	// generate smtpAddress
// 	// output : smtp.gmail.com:587
// 	smtpAddress := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

// 	// proses kirim email
// 	err = smtp.SendMail(smtpAddress, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))

// 	return
// }

func sendMailGoMail(to []string, cc []string, subject string, message string) (err error) {
    // setup gomail message
	mailer := gomail.NewMessage()
    // setting header from 
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
    // setting header to
	mailer.SetHeader("To", to...)

    // setting header CC
	for _, ccEmail := range cc {
		mailer.SetAddressHeader("Cc", ccEmail, "")
	}

    // setting subject
	mailer.SetHeader("Subject", subject)
    // setting body
    // kali ini, kita akan menggunakan body HTML agar tampilan dari emailnya lebih menarik
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)
	return
}

