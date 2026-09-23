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
// Создай массив [3]int со значениями 200, 404, 500.
// Замени вторую ячейку на 201 и верни её.
func task1() int {
	arr := [3]int{200, 404, 500}
	arr[1] = 201

	return arr[1]
}

// Задание 2
// Создай массив [3]int со значениями 10, 20, 30.
// Верни сумму всех ячеек.
func task2() int {
	arr := [3]int{10, 20, 30}
	sum := 0

	for _, v := range arr {
		sum += v
	}

	return sum
}

// Задание 3
// Верни длину массива [4]int. Значения ячеек не важны.
func task3() int {
	return len([4]int{})
}
