package main

import (
	"fmt"
	"math/rand/v2"
)
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main () {
	pass := generatePassword(25)
	fmt.Println(pass, len(pass))
}

func generatePassword(n int) string {
	
	var pass = make([]rune, n)
	for i := range n {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]

	}
	return string(pass)
}