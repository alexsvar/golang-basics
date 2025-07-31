package main

import (
	"demo/app-4/account"
	"demo/app-4/files"
	"demo/app-4/output"
	"fmt"
)

func main() {
	output.PrintError(1)
	fmt.Println("___МЕНЕДЖЕР ПАРОЛЕЙ___")
	vault := account.NewVault(files.NewJsonDb("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("https://a1x.com"))
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1 - Создать аккаунт")
	fmt.Println("2 - Найти аккаунт")
	fmt.Println("3 - Удалить аккаунт")
	fmt.Println("4 - Выход")
	fmt.Scan(&variant)
	return variant
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Login")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url для поиска ")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		fmt.Println("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите url для поиска ")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		fmt.Println("Удалено")
	} else {
		output.PrintError("Не удалось найти аккаунт")
	}

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}
