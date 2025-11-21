package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type account struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Url string `json:"url"`
}

func (acc account) PrintAccount() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *account) generatePassword(n int)  {
	var pass = make([]rune, n)
	for i := range n {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]

	}
	acc.Url = string(pass)
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
		Login: login,
		Password: password,
		Url: urlValue,
		}
	if password == "" {
		acc.generatePassword(15)
	}
	return acc, nil
	}


func ToBytes (acc *account) []byte {
	file, err := json.Marshal(acc) 
	if err != nil {
		fmt.Println(err)
	}
	return file
}