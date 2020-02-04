package helpers

import (
	"github.com/xhit/go-simple-mail"
	"log"
	"time"
	"gokit-example/account/config"
	"html/template"
	"bytes"
	// "fmt"
	// "os"
)

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	// body = buf.String()
	return buf.String(), nil
}

func MailCenter(to string, name string, pathToHtml string,token string) bool {
	conf := config.New()
	templateData := struct {
		Name   string
		Token  string
	}{
		Name: name,
		Token:  conf.AppConfig.ApiUrl + "/account/v1/verify/" + token + "/" + to,
	}
	template, err := ParseTemplate(pathToHtml, templateData)
	server := mail.NewSMTPClient()
	server.Host = conf.Mailing.Host
	// TODO better solution
	server.Port = 587
	server.Username =  conf.Mailing.Username
	server.Password = conf.Mailing.Password
	server.Encryption = mail.EncryptionTLS
	
	//Variable to keep alive connection
	server.KeepAlive = false
	
	//Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second
	
	//Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second
	
	//SMTP client
	smtpClient,err :=server.Connect()
	
	if err != nil{
		log.Fatal(err)
	}

	//New email simple html with inline and CC
	email := mail.NewMSG()

	email.SetFrom(server.Username).
		AddTo(to).
		SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, template)

	//Call Send and pass the client
	err = email.Send(smtpClient)

	if err != nil {
		log.Println(err)
		return false
	} else {
		log.Println("Email Sent")
		return true
	}
	// return true
}