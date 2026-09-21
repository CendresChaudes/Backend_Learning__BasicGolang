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
// Создай переменную city со значением "Казань".
// Верни значение переменной.
func task1() string {
	city := "Казань"
	return city
}

// Задание 2
// Создай переменную balance со значением 100.
// Потом положи в неё 250 и верни новое значение.
func task2() int {
	balance := 100
	balance = 250
	return balance
}

// Задание 3
// Создай две переменные: firstName "Анна" и lastName "Иванова".
// Верни их одной строкой через пробел.
func task3() string {
	firstName := "Анна"
	lastName := "Иванова"
	return firstName + " " + lastName
}
