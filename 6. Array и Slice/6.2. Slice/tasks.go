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
// Создай срез []int{200, 404}. Допиши 500 через append.
// Верни новую длину.
func task1() int {
	slice := []int{200, 404}
	slice = append(slice, 500)

	return len(slice)
}

// Задание 2
// Создай срез через make([]int, 2, 4). Верни его cap.
func task2() int {
	slice := make([]int, 2, 4)
	return cap(slice)
}

// Задание 3
// Возьми окно [1:3] у среза []int{200, 404, 500, 201}.
// Верни первый элемент окна.
func task3() int {
	oldSlice := []int{200, 404, 500, 201}
	newSlice := oldSlice[1:3]

	return newSlice[0]
}
