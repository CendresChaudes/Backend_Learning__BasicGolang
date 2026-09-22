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
// Опиши func total(nums ...int) int.
// Сложи все переданные числа. Если чисел нет, сумма равна 0.
// В task1 верни total(10, 20, 12).
func task1() int {
	return 0
}

// Задание 2
// Опиши func line(prefix string, parts ...string) string.
// Начни текст с prefix и допиши к нему все parts по порядку.
// В task2 верни line("user:", "an", "na").
func task2() string {
	return ""
}

// Задание 3
// Опиши func count(nums ...int) int. Она возвращает len(nums).
// В task3 создай срез []int{7, 8, 9}.
// Верни count от этого среза: после имени среза поставь ...
func task3() int {
	return 0
}
