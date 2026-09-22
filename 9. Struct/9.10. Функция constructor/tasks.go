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
// На уровне пакета опиши type User struct с полями Name string и Age int.
// Функция NewUser(name string, age int) User заполняет оба поля из аргументов.
// Верни Name у NewUser("Ian", 33).
func task1() string {
	return ""
}

// Задание 2
// На уровне пакета опиши type Account struct с полями Owner string и Balance int.
// Функция NewAccount(owner string) Account ставит Owner из аргумента и Balance 100.
// Верни Balance у NewAccount("Mia").
func task2() int {
	return 0
}

// Задание 3
// На уровне пакета опиши type Point struct с полями X int и Y int.
// Функция NewPoint(x int, y int) Point заполняет оба поля.
// Верни сумму X и Y у NewPoint(2, 8).
func task3() int {
	return 0
}
