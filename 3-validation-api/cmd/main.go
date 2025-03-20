package main

import (
	"fmt"
	"go-training/validation-api/configs"
	"go-training/validation-api/internal/mailer"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	mailer.NewMailerHttpHandler(router, mailer.MailerDeps{
		SmtpConfig:       conf.SmtpConfig,
		SendEmailHandler: *mailer.NewSendEmailHandler(conf.SmtpConfig),
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on 8081")
	server.ListenAndServe()

}
