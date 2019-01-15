package mx

import (
	"log"
	"net/smtp"
	"os"
)

const recipient = "rozhkov@mc.asu.ru"

var (
	login = os.Getenv("MX_LOGIN")
	pass  = os.Getenv("MX_PASS")
)

// Check verifies working of ASU's mail in mx.asu.ru
func Check() []byte {
	auth := smtp.PlainAuth("", login, pass, "mx.asu.ru")

	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: Mail is working!\r\n" +
		"\r\n" +
		"There is nothing interesting.\r\n")

	err := smtp.SendMail("mx.asu.ru:25", auth, "Checker@asu.ru", to, msg)
	if err != nil {
		log.Println("unable to send email. ", err)
		return []byte("false")
	}

	log.Println("Почта на mx.asu.ru успешно работает!")
	return []byte("true")
}
