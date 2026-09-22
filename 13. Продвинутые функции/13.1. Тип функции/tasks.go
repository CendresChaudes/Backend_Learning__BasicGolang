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
// Опиши type PortCheck func(int) bool.
// Опиши func isWeb(port int) bool. Она возвращает true, если port равен 443.
// В task1 запиши isWeb в переменную типа PortCheck и верни вызов с аргументом 443.
func task1() bool {
	return false
}

// Задание 2
// Опиши type Label func(string) string.
// Опиши func vaultLabel(name string) string. Она возвращает "vault:" + name.
// В task2 запиши vaultLabel в переменную типа Label и верни вызов с "mail".
func task2() string {
	return ""
}

// Задание 3
// Опиши func add(a, b int) int. Она возвращает сумму a и b.
// В task3 запиши add в переменную типа func(int, int) int.
// Верни вызов этой переменной с аргументами 20 и 22.
func task3() int {
	return 0
}
