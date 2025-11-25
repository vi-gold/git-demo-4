package main

import "fmt"

// 9. Struct
type account struct {
	login    string
	password string
	url      string
}

func main() {
	str := []rune("Привет!)")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(&myAccount)
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

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
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
