// https://kovardin.ru/articles/go/smtp-server/
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"slices"
	"strings"
	//"os"

	"github.com/emersion/go-smtp"
)

var addr = ":25"

var allowlistTo = ""

func init() {
	flag.StringVar(&addr, "l", addr, "Listen address")
	flag.StringVar(&allowlistTo, "a", "", "allowed receiving email addresses")
}

type backend struct{}

func (bkd *backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &session{}, nil
}

type session struct {
	From string
	To   []string
}

func (s *session) AuthPlain(username, password string) error {
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	//fmt.Println("Mail from:", from)
	s.From = from
	return nil
}

func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	fmt.Println("Rcpt to:", to)
	s.To = append(s.To, to)
	return nil
}

func (s *session) Data(r io.Reader) error {
	if b, err := io.ReadAll(r); err != nil {
		return err
	} else {
		if slices.IndexFunc(s.To, func(str string) bool {
			return strings.HasSuffix(str, allowlistTo)
		}) > -1 {
			fmt.Println("Received message:", string(b))
		}
		// Тут происходит обработка письма
		return nil
	}
}

func (s *session) Reset() {}

func (s *session) Logout() error {
	return nil
}

func main() {
	flag.Parse()

	s := smtp.NewServer(&backend{})

	s.Addr = addr
	s.Domain = "localhost"
	s.AllowInsecureAuth = true
	//s.Debug = os.Stdout

	log.Println("Starting SMTP server at", addr)
	log.Fatal(s.ListenAndServe())
}
