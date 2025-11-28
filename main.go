package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println(rand.IntN(10))

	str := []rune("Привет!)")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}
	myAccount.outputPassword()
	fmt.Println(myAccount)
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
