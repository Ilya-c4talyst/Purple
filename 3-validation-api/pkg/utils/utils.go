package utils

import (
	"math/rand"
	"net/smtp"

	"github.com/jordan-wright/email"
)

var Letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateHash() string {
	// В версии 1.24 seed не ставится
	var hash string
	for range 5 {
		hash += string(Letters[rand.Intn(len(Letters))])
	}

	return hash
}

func SendMail(emailTo string, hash string, emailOut string, password string, address string) error {
	e := email.NewEmail()
	e.From = "<" + emailOut + ">"
	e.To = []string{emailTo}
	e.Text = []byte("Перейдите по ссылке для подтверждения: http://localhost:8080/verify/" + hash)
	// EmailOut - отправитель, password - пароль SMTP, address - SMTP address
	err := e.Send(address+":587", smtp.PlainAuth("", emailOut, password, address))

	if err != nil {
		return err
	}
	return nil

}

func Remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
