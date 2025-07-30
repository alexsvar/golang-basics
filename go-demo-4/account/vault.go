package account

import (
	"demo/app-4/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Struct Vault
type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Функция-конструктор Vault
func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

// Добавление аккаунта с преобразованием
func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}

// Преобразование
func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Поиск аккаунта по URL
func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

// Удаление аккаунта по URL
func (vault *Vault) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
	return isDeleted
}
