
package main

// https://gist.github.com/chrisgillis/10888032
// https://gist.github.com/andelf/5004821


import (
  "log"
	"net/mail"
	"encoding/base64"
	"net/smtp"
	"fmt"
	"strings"
)

func encodeRFC2047(String string) string{
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}


func main() {
	// Set up authentication information.

	smtpServer := "smtp.rsnweb.ch"
	
	// https://golang.org/pkg/net/smtp/#PlainAuth
	auth := smtp.PlainAuth(
		"",
		"username@example.com",
		"TOP_SECRET",
		smtpServer,
	)

	from := mail.Address{"监控中心", "username@example.com"}
	to := mail.Address{"收件人", "username@domain.com"}
	title := "This is a test当前时段统计报表"

	body := "Some text... 报表内容一切正常";

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer + ":25",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
		//[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
