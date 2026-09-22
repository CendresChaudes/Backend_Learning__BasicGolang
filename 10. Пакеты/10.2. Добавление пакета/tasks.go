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
// Создай пакет wallet: папка wallet и файл с первой строкой package wallet.
// Функция Owner() string возвращает "Mia".
// Импорт: "lesson10/add/wallet".
// Верни результат wallet.Owner().
func task1() string {
	return ""
}

// Задание 2
// В том же пакете wallet добавь функцию Balance() int.
// Она возвращает 100.
// Верни результат wallet.Balance().
func task2() int {
	return 0
}

// Задание 3
// В том же пакете wallet добавь функцию Label(owner string) string.
// Она возвращает "wallet: " и сразу имя владельца.
// Верни wallet.Label("Mia").
func task3() string {
	return ""
}
