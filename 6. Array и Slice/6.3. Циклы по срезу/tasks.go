package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1([]int{200, 404, 500}))

	fmt.Println("Задание 2")
	fmt.Println(task2([]int{100, 404, 200, 500}))

	fmt.Println("Задание 3")
	fmt.Println(task3([]int{200, 201, 404}, 404))
}

// Задание 1
// Верни сумму codes. Обходи срез через range.
func task1(codes []int) int {
	var sum int
	for _, v := range codes {
		sum += v
	}

	return sum
}

// Задание 2
// Верни, сколько чисел в codes строго больше 300.
func task2(codes []int) int {
	var count int
	for _, v := range codes {
		if v > 300 {
			count++
		}
	}
	return count
}

// Задание 3
// Верни индекс первого элемента, равного target.
// Если такого элемента нет, верни -1.
func task3(codes []int, target int) int {
	for i, v := range codes {
		if v == target {
			return i
		}
	}

	return -1
}
