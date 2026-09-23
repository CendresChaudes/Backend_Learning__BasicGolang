package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(1, 2, 3))

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3([]int{4, 5}))
}

// Задание 1
// Опиши func total(nums ...int) int.
// Верни сумму аргументов.
// task1 вызывает total с переданными числами.
func task1(nums ...int) int {
	return 0
}

// Задание 2
// Вызови total() и total(10). Верни сумму этих двух результатов.
// Пустой вызов даёт 0, поэтому ответ равен результату total(10).
func task2() int {
	return 0
}

// Задание 3
// Вызови total, раскрыв срез nums через ... . Верни сумму.
func task3(nums []int) int {
	return 0
}
