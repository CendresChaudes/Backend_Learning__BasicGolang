package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1([]int{0, 1, 2, 3, 4}, 2))

	names := []string{"выход", "показать", "добавить", "найти", "удалить"}
	fmt.Println("Задание 2")
	fmt.Println(task2(names, "добавить"))

	fmt.Println("Задание 3")
	fmt.Println(task3(names, "найти"))
}

// Задание 1
// Опиши func Index[T comparable](items []T, answer T) int.
// Верни индекс первого элемента, равного answer.
// Если такого элемента нет, верни -1.
// task1 вызывает Index и возвращает его результат.
func task1(items []int, answer int) int {
	return 0
}

// Задание 2
// Опиши func Allowed[T comparable](items []T, answer T) bool.
// Верни true, если answer равен одному из элементов.
// Если совпадения нет, верни false.
// task2 вызывает Allowed и возвращает его результат.
func task2(items []string, answer string) bool {
	return false
}

// Задание 3
// Опиши func Code[T comparable](items []T, answer T) string.
// Если answer есть в items, верни "ok". Иначе верни "нет".
// Для проверки можно вызвать Allowed.
// task3 вызывает Code и возвращает его результат.
func task3(items []string, answer string) string {
	return ""
}
