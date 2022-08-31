package main

import (
	"flag"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Post struct {
	Name    string
	Address string
	Subject string
	Content string
}

func main() {
	url := flag.String("url", "foo", "a string")
	token := flag.String("token", "foo", "a string")
	revicer := flag.String("revicer", "foo", "a string")
	title := flag.String("title", "foo", "a string")
	body := flag.String("body", "foo", "a string")
	flag.Parse()

	log.Println(url, token)

	pop3Url := "smtp.qq.com:587"
	auKey := "youraukey"
	name := "user"

	var receiverName []string
	var receiverAddr []string

	receiverName = append(receiverName, name)
	receiverAddr = append(receiverAddr, *revicer)

	MailWithToken(pop3Url, auKey, receiverName, *title, receiverAddr, []byte(*body))
}

func MailWithToken(pop3Url string, auKey string, receiverName []string, subject string, receiverAddr []string, content []byte) {
	log.Println(pop3Url, auKey, receiverName, subject, receiverAddr, content)

	em := email.NewEmail()
	em.From = "402513681@qq.com"
	em.To = receiverAddr

	em.Subject = subject
	em.Text = content

	err := em.Send(pop3Url, smtp.PlainAuth("", em.From, auKey, "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("send successfully ... ")
}
