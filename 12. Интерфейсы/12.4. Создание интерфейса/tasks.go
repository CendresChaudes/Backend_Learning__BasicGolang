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
// На уровне пакета опиши интерфейс Labeler с методом Label() string.
// Опиши type Disk struct с полем Name string.
// Метод Label у Disk возвращает "disk:" и сразу за ней Name.
// Опиши функцию show(item Labeler) string. Она возвращает item.Label().
// В task1 верни show от Disk с именем "vault".
func task1() string {
	return ""
}

// Задание 2
// Опиши type Cloud struct с полем Name string.
// Метод Label у Cloud возвращает "cloud:" и сразу за ней Name.
// В task2 верни show от Cloud с именем "vault".
func task2() string {
	return ""
}

// Задание 3
// Опиши интерфейс Checker с методом OK() bool.
// Метод OK у Disk возвращает true, если Name не пустой.
// Опиши функцию opened(item Checker) bool. Она возвращает item.OK().
// В task3 верни opened от Disk с именем "vault".
func task3() bool {
	return false
}
