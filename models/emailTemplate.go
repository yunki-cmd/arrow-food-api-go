package models

import (
	"fmt"
)

type EmailTemplate struct {
	Sender  string
	Subject string
	Body    string
}

func (e *EmailTemplate) BuildTemplate() string{
	header := make(map[string]string)
	header["From"] = e.Sender
	header["Subject"] = e.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "quoted-printable"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + e.Body
	fmt.Println("mensaje: ", message)
	return message
}