package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRune = []rune("qwertyuiopasdfghjklzxcvbnm0123456789!@#$%^&-")

// 9. Struct
type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	cratedAt  time.Time
	updatedAt time.Time
	account
}

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(length int) {
	pass := make([]rune, length)
	for i := range pass {
		pass[i] = letterRune[rand.IntN(len(letterRune))]
	}
	acc.password = string(pass)
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	if err != nil {
		return nil, errors.New("IVALID_URL")
	}
	newAcc := &accountWithTimeStamp{
		cratedAt:  time.Now(),
		updatedAt: time.Now(),
		account: account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if newAcc.password == "" {
		var length int
		fmt.Print("\nВведите длинну пароля: ")
		fmt.Scan(&length)
		newAcc.account.generatePassword(length)
	}
	return newAcc, nil
}

// func newAccount(login, password, urlString string) (*account, error) {
// 	_, err := url.ParseRequestURI(urlString)
// 	if login == "" {
// 		return nil, errors.New("EMPTY_LOGIN")
// 	}
// 	if err != nil {
// 		return nil, errors.New("IVALID_URL")
// 	}
// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}
// 	if newAcc.password == "" {
// 		var length int
// 		fmt.Print("\nВведите длинну пароля: ")
// 		fmt.Scan(&length)
// 		newAcc.generatePassword(length)
// 	}
// 	return newAcc, nil
// }
