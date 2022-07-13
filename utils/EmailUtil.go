package utils

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

func SendEmail(invitation string, to string) {
	from := "no-replay@alphacorp.vsk.gr"
	fmt.Println("sending an email to: " + to)
	fmt.Println("invitation: ")

	// todo move credentials to .env
	password := "BEc+pE1TZj/tLQJLq3yxO3neZk020+QvqXUm3KYekMqu"
	smtpHost := "email-smtp.eu-central-1.amazonaws.com"
	smtpPort := "587"

	// improve template add name surname ...
	body := "Welcome to Alphacorp to activate your account please click the invitation link http://localhost:8080/register/" + invitation

	// todo improve headers
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = "🔏 Alphacorp account registration invitation"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	auth := smtp.PlainAuth("", "AKIAW66CCHMX64VN5JD3", password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Send successfully")

}
