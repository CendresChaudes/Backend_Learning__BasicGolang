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
// Верни команду, которая приводит go.mod в соответствие с импортами кода.
func task1() string {
	return ""
}

// Задание 2
// Верни имя файла рядом с go.mod, в котором лежат контрольные суммы модулей.
func task2() string {
	return ""
}

// Задание 3
// Код импортирует пакет, а строки require для него в go.mod ещё нет.
// Верни true, если go mod tidy добавит эту зависимость.
func task3() bool {
	return false
}
