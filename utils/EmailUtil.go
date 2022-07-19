package utils

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

func SendEmail(invitation string, to string) error {
	from := Config.EmailVars.From
	host := Config.HostDomain
	protocol := "https"
	if host == "localhost" {
		host = host + ":8080"
		protocol = "http"
	}

	smtpPort := "587"
	// todo change url depending on the .env value
	body := "👋 Welcome to Alphacorp❗️\n " +
		"to activate your account please click the invitation link: \n" +
		protocol + "://" + host + "/register/" + invitation + "\n" +
		"Follow the instructions, download the app on your phone 📱 and get some credentials. 📜 \n \n" +
		"*Ps: 'Please do not share this email' \n \n" +
		"Alphacorp access management administration"

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
	auth := smtp.PlainAuth("", Config.EmailVars.Username, Config.EmailVars.Password, Config.EmailVars.Host)
	err := smtp.SendMail(Config.EmailVars.Host+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	fmt.Println("Send successfully")
	return nil
}
