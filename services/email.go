package services

import (
	utils "Modifa2/utils"

	"gopkg.in/gomail.v2"
)

/*Send ... */
func Send(To string, Message string) {

	from := utils.MustGetenv("EmailFrom")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", To)
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Proect Main")
	m.SetBody("text/html", Message)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "projectcommunication123@gmail.com", "Modifa1993*")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
