package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(2))

	fmt.Println("Задание 2")
	fmt.Println(task2(2))

	fmt.Println("Задание 3")
	fmt.Println(task3([]int{2, 3, 0}))
}

// Задание 1
// Верни имя команды.
// 0 — "выход", 1 — "показать", 2 — "добавить", 3 — "найти", 4 — "удалить".
// Любой другой код — "неизвестно".
func task1(code int) string {
	return ""
}

// Задание 2
// Верни true, если code — известная команда: 0, 1, 2, 3 или 4.
func task2(code int) bool {
	return false
}

// Задание 3
// Пройди codes слева направо.
// Ноль прерывает проход и сам в ответ не входит.
// Если нуля нет, возьми все коды.
// Имена команд до остановки склей через пробел.
// Имя бери той же расшифровкой, что в задании 1.
func task3(codes []int) string {
	return ""
}
