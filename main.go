package main

import (
	"fmt"
	"math/rand/v2"
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

func main () {
	login := promptData("Введите логин")
	url := promptData("Введите url")

	myAcc := account{
		login: login,
		url: url,
	}

	myAcc.generatePassword(15)
	myAcc.printAccount()

}

func promptData (message string) string {
	var res string
	fmt.Println(message)
	fmt.Scan(&res)
	return res
}