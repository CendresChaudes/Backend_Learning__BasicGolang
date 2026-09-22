package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(100, []int{50, -30, -20}))

	fmt.Println("Задание 2")
	fmt.Println(task2([]int{-10, 5, -3}))

	fmt.Println("Задание 3")
	fmt.Println(task3([]int{10, -4}))
}

// Задание 1
// Верни баланс после всех операций. Начни со start и прибавляй ops по очереди.
func task1(start int, ops []int) int {
	return 0
}

// Задание 2
// Верни сумму расходов положительным числом. Приходы не учитывай.
func task2(ops []int) int {
	return 0
}

// Задание 3
// Старт равен 0. После каждой операции запоминай баланс.
// Верни наименьший из этих балансов. Сам старт 0 не учитывай.
// Если операций нет, верни 0.
func task3(ops []int) int {
	return 0
}
