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
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CratedAt  time.Time `json:"cratedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.Login)
	fmt.Println(acc.Password, acc.Url)
}

func (acc *Account) generatePassword(length int) {
	pass := make([]rune, length)
	for i := range pass {
		pass[i] = letterRune[rand.IntN(len(letterRune))]
	}
	acc.Password = string(pass)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	if err != nil {
		return nil, errors.New("IVALID_URL")
	}
	newAcc := &Account{
		CratedAt:  time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}

	// Демонстрация тегов в полях структуры для json
	//field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	//fmt.Println(string(field.Tag))

	if newAcc.Password == "" {
		var length int
		fmt.Print("\nВведите длинну пароля: ")
		fmt.Scan(&length)
		newAcc.generatePassword(length)
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
