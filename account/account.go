package account

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc account) PrintAccount() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *account) generatePassword(n int) {
	var pass = make([]rune, n)
	for i := range n {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(pass)
}

func NewAccount(scanner *bufio.Scanner, storage *Storage) (error) {
	login := promptDataWithScanner(scanner, "Введите логин")
	password := promptDataWithScanner(scanner, "Введите пароль (оставьте пустым для автогенерации)")
	urlValue := promptDataWithScanner(scanner, "Введите url")

	_, err := url.ParseRequestURI(urlValue)
	if err != nil {
		fmt.Println("неверный URL: %w", err)
		return err
	}

	acc := &account{
		Login:    login,
		Password: password,
		Url:      urlValue,
		CreatedAt: time.Now(),
    	UpdatedAt: time.Now(),

	}

	if password == "" {
		acc.generatePassword(15)
		fmt.Printf("Сгенерирован пароль: %s\n", acc.Password)
	}

	storage.Accounts = append(storage.Accounts, *acc)

	err = storage.Save()
	if err != nil {
		return err
	}
	return nil
}


func promptDataWithScanner(scanner *bufio.Scanner, message string) string {
	fmt.Println(message)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}