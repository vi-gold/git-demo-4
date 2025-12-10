package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	//files.ReadFile()
	// Обращаемся к нашему пакету "demo/password/files"
	// В нем вызываем функцию записи и передаем туда
	// что записать и имя файла куда
	//files.WriteFile("Привет!!! Я файл", "file.txt")

	//Tests
	// fmt.Println(rand.IntN(10))
	// str := []rune("Привет!)")
	// for _, ch := range string(str) {
	// 	fmt.Println(ch, string(ch))
	// }

	// 8.4 Использование указателей
	// a := 5
	// double(&a)
	// fmt.Println(a)

	// 8.5 Reverse массива
	// a := [4]int{1, 2, 3, 4}
	// reverse(&a) // Меняет порядок элементов на обратный
	// fmt.Println(a)

Menu:
	for {
		operation := getMenu()
		switch operation {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		case 4:
			fmt.Println("Выход")
			break Menu
		}

	}
}

func getMenu() int {
	fmt.Print("---Менеджер паролей---\n-Меню-\n1. Создать аккаунт\n2. Найти аккаунт\n3. Удалить аккаунт\n4. Выход\nВведите операцию: ")
	var operation int
	fmt.Scanln(&operation)
	if operation < 1 || operation > 4 {
		color.Cyan("Не валидная операция.\nВведите повторно!")
		fmt.Println()
	}
	return operation
}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}
	myAccount.OutputPassword()
	//fmt.Println(myAccount)

	// Format newAccont to byte before create []vaultAccounts
	//file, err := myAccount.ToBytes()
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать данные в json")
		return
	}
	files.WriteFile(data, "data.json")
}

func findAccount() {
	// TODO
}
func deleteAccount() {
	//TODO
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
