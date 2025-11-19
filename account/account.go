package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type account struct {
	login string
	password string
	url string
}

func (acc account) PrintAccount() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int)  {
	var pass = make([]rune, n)
	for i := range n {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]

	}
	acc.password = string(pass)
}

func NewAccount(login, password, urlValue string) (*account, error) {

	if login == "" {
		return nil, errors.New("invalid login")
	}
	_, err := url.ParseRequestURI(urlValue)
	if err != nil {
		return nil, errors.New("Invalid url")
	}

	acc := &account{
		login: login,
		password: password,
		url: urlValue,
		}
	if password == "" {
		acc.generatePassword(15)
	}
	return acc, nil
	}
