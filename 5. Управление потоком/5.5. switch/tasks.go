package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("POST"))

	fmt.Println("Задание 2")
	fmt.Println(task2(201))

	fmt.Println("Задание 3")
	fmt.Println(task3(6))
}

// Задание 1
// Через switch: "GET" → "чтение", "POST" → "создание".
// Для любой другой строки верни "неизвестно".
func task1(method string) string {
	return ""
}

// Задание 2
// Через switch: 200 и 201 → "успех", 404 → "нет".
// Для остальных чисел верни "другое".
func task2(code int) string {
	return ""
}

// Задание 3
// Через switch: 1, 2, 3, 4 и 5 → "будни".
// 6 и 7 → "выходные". Для остальных чисел верни "нет".
func task3(day int) string {
	return ""
}
