package main

import "fmt"

type account struct {
	login string
	password string
	url string
}

func main() {
	str := []rune("Привет!)")
	for _, char := range string(str) {
		fmt.Println(char, string(char))
	}

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount := account{
		login: login,
		password: password,
		url: url,
	}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}

func outputPassword(acc *account) {
	fmt.Println(acc)
	fmt.Println((*acc).login, acc.password, acc.url)
}