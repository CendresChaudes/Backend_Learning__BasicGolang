package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1([]int{1, 2, 3, 4}))

	fmt.Println("Задание 2")
	fmt.Println(task2([]int{200, 404, 500}))

	fmt.Println("Задание 3")
	fmt.Println(task3([]int{200, 404}))
}

// Задание 1
// Верни сумму чисел. Обойди срез через range.
func task1(nums []int) int {
	return 0
}

// Задание 2
// Верни индекс первого числа 404. Если такого нет, верни -1.
func task2(codes []int) int {
	return 0
}

// Задание 3
// Собери коды в одну строку через дефис: 200 и 404 дадут "200-404".
// В начале и в конце дефиса нет. Обойди срез через range.
func task3(codes []int) string {
	return ""
}
