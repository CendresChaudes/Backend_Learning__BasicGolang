package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(404))

	fmt.Println("Задание 2")
	fmt.Println(task2(70))

	fmt.Println("Задание 3")
	fmt.Println(task3(-2))
}

// Задание 1
// 200 → "ok", 404 → "нет". Для остальных чисел верни "другое".
func task1(code int) string {
	return ""
}

// Задание 2
// score от 90 и выше → "отлично".
// score от 60 и выше → "зачёт".
// Ниже 60 → "незачёт".
func task2(score int) string {
	return ""
}

// Задание 3
// Положительное n → "плюс". Отрицательное → "минус". Ноль → "ноль".
func task3(n int) string {
	return ""
}
