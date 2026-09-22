package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(10, 20, 30))

	fmt.Println("Задание 2")
	fmt.Println(task2([]int{10, 20}, []int{30, 40}))

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Параметр nums собирает сколько угодно чисел. Верни их сумму.
func task1(nums ...int) int {
	return 0
}

// Задание 2
// Добавь к base все числа из extra одним append. extra передай через ....
// Верни сумму получившегося среза.
func task2(base []int, extra []int) int {
	return 0
}

// Задание 3
// Объяви функцию с параметром ...int, которая возвращает len этих аргументов.
// Вызови её, распаковав срез []int{8, 8, 8, 8}. Верни это количество.
func task3() int {
	return 0
}
