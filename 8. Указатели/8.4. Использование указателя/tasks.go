package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	n := 10
	fmt.Println(task2(&n))

	fmt.Println("Задание 3")
	a, b := 2, 9
	fmt.Println(task3(&a, &b))
}

// Задание 1
// Создай n := 3 и указатель p на n.
// Запиши через p число 10. Верни n.
func task1() int {
	return 0
}

// Задание 2
// Прибавь 5 к числу, на которое указывает n.
// Запиши сумму обратно по этому указателю и верни её.
func task2(n *int) int {
	return 0
}

// Задание 3
// Поменяй местами числа, на которые указывают a и b.
// Верни число, которое после обмена лежит по указателю a.
func task3(a, b *int) int {
	return 0
}
