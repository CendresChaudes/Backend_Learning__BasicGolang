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
// Создай анонимную функцию func(n int) int, которая возвращает n * 3.
// Запиши её в переменную, вызови с аргументом 4 и верни результат.
func task1() int {
	return 0
}

// Задание 2
// Вызови анонимную функцию сразу: скобки с аргументом стоят после тела.
// Функция принимает string и возвращает "login:" + это имя.
// Аргумент вызова — "root". Верни результат.
func task2() string {
	return ""
}

// Задание 3
// Опиши func keep(n int, ok func(int) bool) bool. Она возвращает ok(n).
// В task3 верни keep(443, ...).
// Вторым аргументом передай анонимную функцию:
// она возвращает true, если число равно 443.
func task3() bool {
	return false
}
