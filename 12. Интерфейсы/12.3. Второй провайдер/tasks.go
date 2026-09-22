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
// На уровне пакета опиши type Disk struct с полем Name string.
// Метод Label возвращает "disk:" и сразу за ней Name.
// Создай Disk с именем "vault" и верни Label.
func task1() string {
	return ""
}

// Задание 2
// На уровне пакета опиши type Cloud struct с полем Name string.
// Метод Label возвращает "cloud:" и сразу за ней Name.
// Создай Cloud с именем "vault" и верни Label.
func task2() string {
	return ""
}

// Задание 3
// У Cloud метод Ready возвращает true, если Name не пустой.
// Создай Cloud с именем "vault" и верни Ready.
func task3() bool {
	return false
}
