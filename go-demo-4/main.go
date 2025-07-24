package main

import "fmt"

type account struct {
	login string
	password string
	url string
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	// account1 := account{
	// 	login,
	// 	password,
	// 	url,
	// }
	account1 := account{
		login: login,
		password: password,
		url: url,
	}

	outputPassword(login, password, url)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}

func outputPassword(login, password, url string) {
	fmt.Println(login, password, url)
}