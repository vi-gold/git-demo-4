package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна!")
}

func ReadFileTest() {
	// Для открытия и чтения по байтам
	// file, err := os.Open("files.txt")

	data, err := os.ReadFile("file.txt")
	// Проверяем на отсутствие ошибок при чтении файла
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}

// Принимаем аргументы из main в функцию записи
func WriteFileTest(content string, name string) {
	// Создаем файл с полученным именем
	// Функция создания возвращает указатель на структуру
	// файла И ошибку при ее наличии
	file, err := os.Create(name)

	// Проверяем на отсутствие ошибок при создании файла
	if err != nil {
		fmt.Println(err)
	}

	// Выполнения записи в файл, функция
	// возвращает длинна записанных байт и ошибку при наличии
	_, err = file.WriteString(content)

	// Закрытия файла после выолнения всего
	// содержимого функции, благодаря
	// ключевому слову defer
	defer file.Close()

	// Проверка на ошибки при записи в файл
	if err != nil {
		fmt.Println(err)
		return
	}
	// Выводим информацию об успешной записи в файл
	fmt.Println("Запись успешна!")
}
