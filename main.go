package main

import (
	"fmt"
	"github.com/m4nsur/pass-generation-lrn/account"
)


func main () {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный url")
		return
	}

	myAcc.PrintAccount()

}

func promptData (message string) string {
	var res string
	fmt.Println(message)
	fmt.Scanln(&res)
	return res
}