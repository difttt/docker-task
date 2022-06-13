package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)


type Post struct {
	Name string 
	Address string 
	Subject string 
	Content string 
}

func main() {
	
	server := http.Server{
		Addr: "localhost:3030",
	}
	//landing page
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
		http.ServeFile(w,r,"form.html")
	})

	fmt.Printf("starting server at port 3030\n")
	//handle request
	http.HandleFunc("/form", HandEmail) 
	
	
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func HandEmail(w http.ResponseWriter, r *http.Request) {
	
	
	switch r.Method {
		case http.MethodPost:
		dec := json.NewDecoder(r.Body)

		post := Post{}

		err := dec.Decode(&post)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}
		var receiverName []string
		var receiverAddr []string
		receiverName = append(receiverName,post.Name)
		receiverAddr = append(receiverAddr,post.Address)
		fmt.Println(receiverName,receiverAddr)
		subject := post.Subject
		content := []byte(post.Content)

		pop3Url := "smtp.qq.com:25"
		auKey := "xxx"
	
		MailWithToken(pop3Url,auKey,receiverName,subject,receiverAddr,content) 
	
	}
}

func MailWithToken(pop3Url string,auKey string,receiverName []string,subject string,receiverAddr []string,content []byte) {

	em := email.NewEmail()
	em.From = "1548546585@qq.com"
	em.To = receiverAddr

	em.Subject = subject
	em.Text = content

	err := em.Send(pop3Url, smtp.PlainAuth("", em.From, auKey, "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("send successfully ... ")
}