package main

import (
	"demo/app-4/account"
	"demo/app-4/files"
	"fmt"
)

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAccount, err := account.NewAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}

	myAccount.OutputPassword()
	files.WriteFile()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}

