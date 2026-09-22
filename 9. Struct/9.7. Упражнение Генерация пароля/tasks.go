package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3("abca", "abc"))
}

// Задание 1
// Алфавит "xyz". Индексы rune: 2, 0, 1.
// Собери пароль из этих символов и верни его.
func task1() string {
	return ""
}

// Задание 2
// Алфавит "ab". Возьми символ с индексом 0 и повтори его 3 раза.
// Верни получившуюся строку.
func task2() string {
	return ""
}

// Задание 3
// Верни true, если каждая rune пароля password есть в алфавите alphabet.
func task3(password, alphabet string) bool {
	return false
}
