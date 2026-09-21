package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3(4, 2.5))
}

// Задание 1
// Создай четыре переменные: count 7 (int), price 2.5 (float64), title "Go" (string), ready true (bool).
// Верни их одной строкой через fmt.Sprintf("%d %.1f %s %t", count, price, title, ready).
// %d — целое, %.1f — дробь с одним знаком, %s — текст, %t — логическое значение.
func task1() string {
	var count = 7
	var price = 2.5
	var title = "Go"
	var ready = true

	return fmt.Sprintf("%d %.1f %s %t", count, price, title, ready)
}

// Задание 2
// Объяви через var четыре переменные тех же типов и ничего им не присваивай.
// Верни их нулевые значения через fmt.Sprintf("%d %.1f %q %t", ...).
// %q покажет пустую строку в кавычках.
func task2() string {
	var count int
	var price float64
	var title string
	var ready bool

	return fmt.Sprintf("%d %.1f %q %t", count, price, title, ready)
}

// Задание 3
// count — число штук (int), price — цена одной штуки (float64).
// Верни сумму. Перед умножением преобразуй count в float64.
func task3(count int, price float64) float64 {
	return float64(count) * price
}
