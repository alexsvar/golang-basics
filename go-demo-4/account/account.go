package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_!?*")

type account struct {
	login string
	password string
	url string
}

type accountWithTimestamp struct{
	createdAt time.Time
	updatedAt time.Time
	acc account
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

func newAccountWithTimestamp(login, password, urlString string) (*accountWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	// Validation
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWithTimestamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		acc: account{
			login: login,
			password: password,
			url: urlString,
		},
	}
	if password == "" {
		newAcc.acc.generatePassword(12)
	}
	return newAcc, nil
}
