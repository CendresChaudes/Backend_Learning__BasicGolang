package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("7"))

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3("Ann"))
}

// Задание 1
// Опиши func withPrefix(prefix string) func(string) string.
// Возвращённая функция склеивает prefix и свой аргумент.
// Верни результат withPrefix("id:")(name).
func task1(name string) string {
	return ""
}

// Задание 2
// Опиши func counter() func() int.
// Возвращённая функция каждый вызов увеличивает своё число на 1 и возвращает его.
// Возьми один счётчик, вызови его два раза и верни второй результат.
func task2() int {
	return 0
}

// Задание 3
// Верни withPrefix("user:")(name).
func task3(name string) string {
	return ""
}
