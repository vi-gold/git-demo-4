package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRune = []rune("qwertyuiopasdfghjklzxcvbnm0123456789!@#$%^&-")

// 9. Struct
type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	cratedAt  time.Time
	updatedAt time.Time
	Account
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc.password, acc.url)
}

func (acc *Account) generatePassword(length int) {
	pass := make([]rune, length)
	for i := range pass {
		pass[i] = letterRune[rand.IntN(len(letterRune))]
	}
	acc.password = string(pass)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	if err != nil {
		return nil, errors.New("IVALID_URL")
	}
	newAcc := &AccountWithTimeStamp{
		cratedAt:  time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if newAcc.password == "" {
		var length int
		fmt.Print("\nВведите длинну пароля: ")
		fmt.Scan(&length)
		newAcc.Account.generatePassword(length)
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
