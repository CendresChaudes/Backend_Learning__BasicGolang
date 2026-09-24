package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(4))

	fmt.Println("Задание 2")
	fmt.Println(task2(3))

	fmt.Println("Задание 3")
	fmt.Println(task3(3))
}

// Задание 1
// Опиши func double(n int) int на уровне пакета.
// Положи её в переменную типа func(int) int и верни вызов с аргументом n.
func double(n int) int {
	return n * 2
}

func task1(n int) int {
	return double(n)
}

// Задание 2
// Опиши func apply(n int, op func(int) int) int.
// Она возвращает op(n).
// Вызови apply с n и функцией double.
func apply(n int, op func(int) int) int {
	return op(n)
}

func task2(n int) int {
	return apply(n, double)
}

// Задание 3
// Вызови apply. Вторым аргументом передай анонимную функцию:
// она возвращает n + 1.
func task3(n int) int {
	return apply(n, func(n int) int {
		return n + 1
	})
}
