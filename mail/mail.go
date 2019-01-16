package mail

import (
	"log"
	"net/smtp"
	"os"
)

const recipient = "rozhkov@mc.asu.ru"

// Mail struct
type Mail struct {
	Name string
}

// Check verifies working of ASU's mail
func (m Mail) Check() []byte {
	var login string
	var pass string
	if m.Name == "mail" {
		login = os.Getenv("MAIL_LOGIN")
		pass = os.Getenv("MAIL_PASS")
	} else {
		login = os.Getenv("MX_LOGIN")
		pass = os.Getenv("MX_PASS")
	}

	auth := smtp.PlainAuth("", login, pass, m.Name+".asu.ru")

	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: Mail is working!\r\n" +
		"\r\n" +
		"There is nothing interesting.\r\n")

	err := smtp.SendMail(m.Name+".asu.ru:25", auth, "Checker@asu.ru", to, msg)
	if err != nil {
		log.Println("unable to send email. ", err)
		return []byte("false")
	}

	log.Println("Почта на " + m.Name + ".asu.ru успешно работает!")
	return []byte("true")
}

// GetName returns name of the mail server
func (m Mail) GetName() string {
	return m.Name
}
