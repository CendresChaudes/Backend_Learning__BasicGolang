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
// Объяви функцию origin, которая возвращает два значения: "Go" и 2009.
// Вызови её из task1 и верни fmt.Sprintf("%s %d", name, year).
func origin() (string, int) {
	return "Go", 2009
}

func task1() string {
	name, year := origin()
	return fmt.Sprintf("%s %d", name, year)
}

// Задание 2
// Объяви функцию split, которая возвращает 10 и 3.
// Вызови её из task2 и верни разность: первое минус второе.

func split() (int, int) {
	return 10, 3
}

func task2() int {
	left, right := split()
	return left - right
}

// Задание 3
// Объяви функцию pair, которая возвращает true и false.
// Вызови её из task3 и верни первое значение.
func pair() (bool, bool) {
	return true, false
}

func task3() bool {
	first, _ := pair()
	return first
}
