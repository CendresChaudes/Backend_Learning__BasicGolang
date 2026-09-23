package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(3))

	fmt.Println("Задание 2")
	fmt.Println(task2(6))

	fmt.Println("Задание 3")
	fmt.Println(task3([]int{1, 4, 15, 2}))
}

// Задание 1
// Верни сумму чисел от 1 до n включительно.
// Если n меньше 1, верни 0.
func task1(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

// Задание 2
// Верни сумму чётных чисел от 1 до n включительно.
// Нечётное число пропускай через continue.
func task2(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			continue
		}

		sum += i
	}

	return sum
}

// Задание 3
// Верни первое число из nums, которое строго больше 10.
// Когда нашёл его, выйди через break.
// Если такого числа нет, верни 0.
func task3(nums []int) int {
	var result int
	for _, num := range nums {
		if num > 10 {
			result = num
			break
		}
	}

	return result
}
