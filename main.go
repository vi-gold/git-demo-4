package main

import (
	"fmt"
	"math/rand/v2"
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

func main() {
	fmt.Println(rand.IntN(10))

	str := []rune("Привет!)")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}

	login := promptData("Введите логин")
	//password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := account{
		login: login,
		//	password: password,
		url: url,
	}

	var length int
	fmt.Print("\nВведите длинну пароля: ")
	fmt.Scan(&length)
	myAccount.generatePassword(length)
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
	fmt.Scan(&res)
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
