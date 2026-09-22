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
// Напиши inner: локальная n равна 10, функция возвращает n.
// Напиши outer: локальная n равна 1, функция возвращает n + inner().
// task1 вызывает outer.
func task1() int {
	return 0
}

// Задание 2
// Напиши step(x int) int. Локальная y равна x + 1, верни y.
// task2 вызывает step(2).
// Наружу уходит копия y. Кадр step к этому моменту уже снят.
func task2() int {
	return 0
}

// Задание 3
// Напиши first, second и third.
// В first локальная n равна 1, верни n + second().
// В second локальная n равна 2, верни n + third().
// В third локальная n равна 4, верни n.
// task3 вызывает first.
func task3() int {
	return 0
}
