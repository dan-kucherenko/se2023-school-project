package emails

import (
	"bufio"
	"os"
)

func GetAllEmails() (emails []string) {
	emailsFileName := os.Getenv("EMAILS_FILE")
	emailsFile, err := os.Open(emailsFileName)
	if err != nil {
		return emails
	}
	defer emailsFile.Close()

	scanner := bufio.NewScanner(emailsFile)
	for scanner.Scan() {
		emails = append(emails, scanner.Text())
	}
	return emails
}

func IsEmailSubscribed(email string) bool {
	isSubscribed := false
	emailsFileName := os.Getenv("EMAILS_FILE")
	emailsFile, _ := os.Open(emailsFileName)
	defer emailsFile.Close()

	scanner := bufio.NewScanner(emailsFile)
	for scanner.Scan() {
		if scanner.Text() == email {
			isSubscribed = true
			break
		}
	}
	return isSubscribed
}

func SubscribeNewEmail(email string) error {
	emailsFileName := os.Getenv("EMAILS_FILE")
	emailsFile, err := os.OpenFile(emailsFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer emailsFile.Close()

	_, err = emailsFile.WriteString(email + "\n")
	return err
}
