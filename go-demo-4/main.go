package main

import "fmt"

type account struct {
	login 		string
	password 	string
	url 			string
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount := account{
		login: 		login,
		password: password,
		url: 			url,
	}

	outputPassword(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}

func outputPassword(acc account) {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}