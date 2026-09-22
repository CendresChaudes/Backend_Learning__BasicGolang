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
// На уровне пакета опиши type Generator struct с полем Alphabet string.
// Метод Password(indexes []int) string собирает пароль:
// каждая rune берётся из Alphabet по очередному индексу.
// Alphabet "xyz", индексы 2, 0, 1. Верни пароль.
func task1() string {
	return ""
}

// Задание 2
// Метод Runes() int возвращает число rune в Alphabet.
// Alphabet "абв". Верни Runes.
func task2() int {
	return 0
}

// Задание 3
// Метод Repeat(n int) string возвращает первую rune алфавита, повторённую n раз.
// Alphabet "xyz", n равен 3. Верни строку.
func task3() string {
	return ""
}
