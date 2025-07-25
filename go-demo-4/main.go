package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login string
	password string
	url string
}

func (acc *account) outputPassword() {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(result)
}

func newAccount(login, password, url string) *account {
	return &account{
		login: login,
		password: password,
		url: url,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_!?*")

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount := newAccount(login, password, url)
	myAccount.generatePassword(12)
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}