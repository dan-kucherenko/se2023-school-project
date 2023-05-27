package emails

import (
	"crypto/tls"
	"fmt"
	"github.com/dan-kucherenko/se-school-project/currency_rate_getter"
	"github.com/go-gomail/gomail"
	"os"
)

// additional function for sending the emails
func sendEmail(emails []string, message string) error {
	if len(emails) == 0 {
		return nil
	}
	senderEmail := os.Getenv("EMAIL")
	senderPassword := os.Getenv("EMAIL_PASSWORD")

	// create a new "message" with headers
	mail := gomail.NewMessage()
	mail.SetHeader("From", senderEmail)
	mail.SetHeader("To", emails...)
	mail.SetHeader("Subject", "Currency rate BTC to UAH")
	mail.SetBody("text/html", message)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPassword)
	// configuration for accepting any certificate presented by the server
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// send the email
	if err := dialer.DialAndSend(mail); err != nil {
		panic(err)
	}
	return dialer.DialAndSend(mail)
}

func SendEmailWithRate() error {
	time, rate, _ := currency_rate_getter.GetRateBtcToUah()
	emails := GetAllEmails()
	message := fmt.Sprintf("At %s, 1 BTC costs %f UAH\n", time, rate)
	return sendEmail(emails, message)
}
