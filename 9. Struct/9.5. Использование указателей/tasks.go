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
// Опиши type Lamp struct с полем Level int.
// Создай значение с Level 1 и указатель на него.
// Через указатель запиши Level 3.
// Верни Level исходной переменной.
func task1() int {
	return 0
}

// Задание 2
// Опиши type Wallet struct с полем Balance int.
// На уровне пакета напиши функцию с параметром *Wallet: она прибавляет 50 к Balance.
// Начальный баланс 100. Верни баланс после вызова.
func task2() int {
	return 0
}

// Задание 3
// Опиши type User struct с полем Name string.
// Создай две переменные: имя "Ann" и имя "Otto".
// Возьми указатель на вторую и через него запиши имя "Eva".
// Верни имя второй переменной.
func task3() string {
	return ""
}
