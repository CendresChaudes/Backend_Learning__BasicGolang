package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	nums3 := [3]int{1, 2, 3}
	fmt.Println(task1(&nums3))

	fmt.Println("Задание 2")
	nums4 := [4]int{10, 20, 30, 40}
	fmt.Println(task2(&nums4))

	fmt.Println("Задание 3")
	nums5 := [5]int{1, 2, 3, 4, 9}
	fmt.Println(task3(&nums5))
}

// Задание 1
// Разверни массив по указателю nums: первый элемент становится последним.
// Верни новый первый элемент.
func task1(nums *[3]int) int {
	return 0
}

// Задание 2
// Разверни массив по указателю nums.
// Верни новый последний элемент.
func task2(nums *[4]int) int {
	return 0
}

// Задание 3
// Разверни массив по указателю nums.
// Верни true, если новый первый элемент больше нового последнего.
func task3(nums *[5]int) bool {
	return false
}
