package main

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"time"
)

type (
	UserAcc struct {
		Username string
		Password string
	}
	Mail struct {
		Sender   UserAcc
		Reciever string
		Body     string
		Subject  string
	}
)

const smtpServer string = "smtp.gmail.com"

var conn *tls.Conn

func b64(inp string) string {
	return base64.StdEncoding.EncodeToString([]byte(inp))
}

func send(text string) {
	log.Println("C: ", text)
	_, err := conn.Write([]byte(text + "\r\n"))

	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
}

func serverRecieve() {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		log.Println("S: ", scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal("err in scanning: ", scanner.Err())
	}
}

func authentication(user UserAcc) {
	send("AUTH LOGIN")
	send(b64(user.Username))
	send(b64(user.Password))
}

func sendMail(mail Mail) {
	send(fmt.Sprintf("MAIL FROM: <%s>", mail.Sender.Username))
	send(fmt.Sprintf("RCPT TO: <%s>", mail.Reciever))
	send("DATA")
	send(fmt.Sprintf("Subject: %s", mail.Subject))
	send(mail.Body)
	send(".")
	send("QUIT")
}

func main() {
	username := flag.String("u", "bob@gmail.com", "your gmail address")
	password := flag.String("p",
		"letmein", "gmail password (you should enable less secure apps")
	reciever := flag.String("dest", "rose@gmail.com",
		"who do you want to send mail to?")
	subject := flag.String("subj", "awesome SMTP client intro", "email subject")
	body := flag.String("body",
		"let me introduce you to newly released email client",
		"email body")

	flag.Parse()

	user := UserAcc{
		Username: *username,
		Password: *password,
	}
	mail := Mail{
		Sender:   user,
		Reciever: *reciever,
		Body:     *body,
		Subject:  *subject,
	}

	conf := &tls.Config{
		//	InsecureSkipVerify = true,
	}
	var err error
	conn, err = tls.Dial("tcp", smtpServer+":465", conf)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go serverRecieve()

	send("EHLO " + smtpServer)
	authentication(user)
	sendMail(mail)

	log.Println("done")
}
