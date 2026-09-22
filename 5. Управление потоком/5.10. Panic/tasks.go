package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2(-3))

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Вызови panic("стоп"). Перехвати его через defer и recover.
// Верни пойманный текст. Дай результату имя, чтобы defer мог его записать.
func task1() string {
	return ""
}

// Задание 2
// Если n меньше 0, вызови panic("минус"), перехвати его и верни -1.
// Если n не меньше 0, верни n. Для такого n panic вызывать не нужно.
func task2(n int) int {
	return 0
}

// Задание 3
// Вызови panic(1). Если recover поймал значение, верни true.
func task3() bool {
	return false
}
