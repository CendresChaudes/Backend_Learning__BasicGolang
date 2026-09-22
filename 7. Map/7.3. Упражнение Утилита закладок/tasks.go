package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(map[string]string{"docs": "https://go.dev"}, "docs"))

	fmt.Println("Задание 2")
	fmt.Println(task2(map[string]string{"docs": "https://go.dev"}, "docs"))

	fmt.Println("Задание 3")
	fmt.Println(task3(map[string]string{"docs": "https://go.dev"}, "blog", "https://go.dev/blog"))
}

// Задание 1
// Верни адрес закладки с именем name.
// Если такого имени нет, верни пустую строку.
func task1(marks map[string]string, name string) string {
	return ""
}

// Задание 2
// Верни true, если закладка с именем name уже есть.
func task2(marks map[string]string, name string) bool {
	return false
}

// Задание 3
// Запиши url под именем name. Если имя уже было, замени адрес.
// Верни число закладок после записи.
func task3(marks map[string]string, name, url string) int {
	return 0
}
