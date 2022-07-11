package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail() {
	from := "no-replay@alphacorp.vsk.gr"
	fmt.Println("sending an email")
	password := "BEc+pE1TZj/tLQJLq3yxO3neZk020+QvqXUm3KYekMqu"
	to := []string{
		"bill1897@yahoo.gr",
	}
	smtpHost := "email-smtp.eu-central-1.amazonaws.com"
	//smtpPort := "587"
	message := []byte("This is a test")

	auth := smtp.PlainAuth("", "AKIAW66CCHMX64VN5JD3", password, smtpHost)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost + ":465",
	}

	conn, err := tls.Dial("tcp", smtpHost+":465", tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Panic(err)
	}

	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to[0]); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write(message)
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

	//err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println("Send successfully")

}
