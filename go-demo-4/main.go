package main

import (
	"demo/app-4/account"
	"demo/app-4/files"
	"fmt"
)

func main() {
	files.WriteFile("It is writing in to file", "file.txt")
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAccount, err := account.NewAccountWithTimestamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}

	myAccount.OutputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}

