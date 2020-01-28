package helpers

import (
	"github.com/xhit/go-simple-mail"
	"log"
	"time"
	"gokit-example/account/config"
)

func MailCenter(to string) bool {
	conf := config.New()
	htmlBody :=
`<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p>Image created by Renee French</p>
	</body>
</html>`

	server := mail.NewSMTPClient()
	
	//SMTP Server
	// server.Host = "smtp.example.com"
	// server.Port = 587
	// server.Username = "test@example.com"
	// server.Password = "examplepass"
	// server.Encryption = mail.EncryptionTLS

	server.Host = conf.Mailing.Host
	// TODO better solution
	server.Port = 587
	server.Username =  conf.Mailing.Username
	server.Password = conf.Mailing.Password
	server.Encryption = mail.EncryptionTLS

	/*
	In version >=2.3.0 you can specified authentication type: PLAIN, LOGIN, CRAM-MD5
	if not specified, default is mail.AuthPlain

	code:
	server.Authentication = mail.AuthPlain
	*/
	
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
		AddTo("test@gmail.com").
		SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)

	// email.AddInline("/path/to/image.png", "Gopher.png")

	//Call Send and pass the client
	err = email.Send(smtpClient)

	if err != nil {
		log.Println(err)
		return false
	} else {
		log.Println("Email Sent")
		return true
	}
}