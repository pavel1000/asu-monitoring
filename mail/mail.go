package mail

import (
	"log"
	"net/smtp"
	"os"
)

const recipient = "rozhkov@vc.asu.ru"

var (
	login = os.Getenv("MAIL_LOGIN")
	pass  = os.Getenv("MAIL_PASS")
)

// Check verifies working of ASU's mail
func Check() bool {
	auth := smtp.PlainAuth("", login, pass, "mx.asu.ru")

	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: Mail is working!\r\n" +
		"\r\n" +
		"There is nothing interesting.\r\n")

	err := smtp.SendMail("mx.asu.ru:25", auth, "Checker@asu.ru", to, msg)
	if err != nil {
		log.Println("unable to send email. ", err)
		return false
	}

	log.Println("Почта успешно работает!")
	return true
}
