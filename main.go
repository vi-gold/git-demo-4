package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var letterRune = []rune("qwertyuiopasdfghjklzxcvbnm0123456789!@#$%^&-")

// 9. Struct
type account struct {
	login    string
	password string
	url      string
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

func newAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	if err != nil {
		return nil, errors.New("IVALID_URL")
	}
	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}
	if newAcc.password == "" {
		var length int
		fmt.Print("\nВведите длинну пароля: ")
		fmt.Scan(&length)
		newAcc.generatePassword(length)
	}
	return newAcc, nil
}

func main() {
	fmt.Println(rand.IntN(10))

	str := []rune("Привет!)")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}
	myAccount.outputPassword()
	// 8.4 Использование указателей
	// a := 5
	// double(&a)
	// fmt.Println(a)

	// 8.5 Reverse массива
	// a := [4]int{1, 2, 3, 4}
	// reverse(&a) // Меняет порядок элементов на обратный
	// fmt.Println(a)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

// 8.5 Reverse массива
// func reverse(a *[4]int) {
// 	fmt.Println(len(a))
// 	for i, value := range *a {
// 		(*a)[len(a)-1-i] = value
// 	}
// }

// 8.4 Использование указателей
// func double(num *int) {
// 	*num = *num * 2
// }
