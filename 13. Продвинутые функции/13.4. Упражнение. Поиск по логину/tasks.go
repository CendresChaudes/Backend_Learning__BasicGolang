package main

import "fmt"

type Account struct {
	Login  string
	Secret string
}

func main() {
	items := []Account{
		{Login: "anna", Secret: "a1"},
		{Login: "boris", Secret: "b2"},
	}

	fmt.Println("Задание 1")
	fmt.Println(task1(items, "boris"))

	fmt.Println("Задание 2")
	fmt.Println(task2(items, "anna"))

	fmt.Println("Задание 3")
	fmt.Println(task3(items, "boris"))
}

// Задание 1
// Опиши func Secret(items []Account, login string, match func(Account, string) bool) string.
// Верни Secret первой записи, для которой match(item, login) равен true.
// Если такой записи нет, верни пустую строку.
// В task1 вызови Secret. Третьим аргументом передай анонимную функцию:
// она возвращает true, когда item.Login равен login.
func task1(items []Account, login string) string {
	return ""
}

// Задание 2
// Опиши func Has(items []Account, login string, match func(Account, string) bool) bool.
// Верни true, если match нашёл запись. Если запись не найдена, верни false.
// В task2 вызови Has с анонимной функцией сравнения логина.
func task2(items []Account, login string) bool {
	return false
}

// Задание 3
// Опиши func Index(items []Account, login string, match func(Account, string) bool) int.
// Верни индекс первой подходящей записи.
// Если записи нет, верни -1.
// В task3 вызови Index с анонимной функцией сравнения логина.
func task3(items []Account, login string) int {
	return 0
}
