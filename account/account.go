package account

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
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

func NewAccount(scanner *bufio.Scanner) ([]byte, error) {
	login := promptDataWithScanner(scanner, "Введите логин")
	password := promptDataWithScanner(scanner, "Введите пароль (оставьте пустым для автогенерации)")
	urlValue := promptDataWithScanner(scanner, "Введите url")
	
	_, err := url.ParseRequestURI(urlValue)
	if err != nil {
		return nil, fmt.Errorf("неверный URL: %w", err)
	}

	acc := &account{
		Login:    login,
		Password: password,
		Url:      urlValue,
	}

	if password == "" {
		acc.generatePassword(15)
		fmt.Printf("Сгенерирован пароль: %s\n", acc.Password)
	}

	file, err := ToBytes(acc)
	if err != nil {
		return nil, fmt.Errorf("ошибка: %w", err)
	}

	return file, nil
}

func ToBytes(acc *account) ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func promptDataWithScanner(scanner *bufio.Scanner, message string) string {
	fmt.Println(message)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}