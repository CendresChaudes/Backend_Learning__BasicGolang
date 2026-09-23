package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(200))

	fmt.Println("Задание 2")
	fmt.Println(task2(200))

	n := 10
	fmt.Println("Задание 3")
	fmt.Println(task3(&n))
}

// Задание 1
// Возьми адрес n и верни число, которое лежит по этому адресу.
func task1(n int) int {
	addr := &n
	return *addr
}

// Задание 2
// Через указатель запиши в переменную n + 1 и верни новое число.
func task2(n int) int {
	addr := &n
	*addr = *addr + 1
	return *addr
}

// Задание 3
// Если n равен nil, верни 0.
// Иначе прибавь к числу по этому адресу 1 и верни новое число.
func task3(n *int) int {
	if n == nil {
		return 0
	}

	*n = *n + 1

	return *n
}
