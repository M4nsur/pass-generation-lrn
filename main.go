package main

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

func (acc account) printAccount() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int)  {
	var pass = make([]rune, n)
	for i := range n {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]

	}
	acc.password = string(pass)
}

func newAccount(login, password, urlValue string) (*account, error) {

	if len([]rune(login)) == 0 {
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
	if len([]rune(password)) == 0 {
		acc.generatePassword(15)
	}
	return acc, nil
	}


func main () {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")

	myAcc, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный url")
		return
	}

	myAcc.printAccount()

}

func promptData (message string) string {
	var res string
	fmt.Println(message)
	fmt.Scanln(&res)
	return res
}