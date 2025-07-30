package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_!?*")

type Account struct {
	login string `json:"login" xml:"test"`
	password string
	url string
}

type AccountWithTimestamp struct{
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(result)
}

func NewAccountWithTimestamp(login, password, urlString string) (*AccountWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	// Validation
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimestamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login: login,
			password: password,
			url: urlString,
		},
	}
	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.Account.generatePassword(12)
	}
	return newAcc, nil
}
