package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// На уровне пакета опиши type Missing struct с полем Path string.
// Метод Error возвращает "нет " и сразу за ней Path.
// Создай Missing с путём "vault.json" и верни Error.
func task1() string {
	return ""
}

// Задание 2
// На уровне пакета опиши type Entry struct с полем Name string.
// Метод String возвращает "entry " и сразу за ней Name.
// Создай Entry с именем "mail" и верни String.
func task2() string {
	return ""
}

// Задание 3
// Опиши функцию text(err error) string.
// Если err равен nil, верни пустую строку.
// Иначе верни err.Error().
// В task3 верни text от Missing с путём "key".
func task3() string {
	return ""
}
