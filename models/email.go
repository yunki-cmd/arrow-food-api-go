package models

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type Email struct {
	Host     string
	Port     string
	Username string
	Password string
	Body     *EmailTemplate
}

func (e *Email) SendMail( body *EmailTemplate) error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	e.Host = "smtp.gmail.com"
	e.Port = "587"
	e.Username = os.Getenv("GMAIL")
	e.Password = os.Getenv("PASSGMAIL")
	e.Body = body

	auth := smtp.PlainAuth("",e.Username,e.Password,e.Host)

	err := smtp.SendMail(e.Host+":"+e.Port, auth,e.Username,[]string{e.Body.Sender},[]byte(e.Body.BuildTemplate()))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
