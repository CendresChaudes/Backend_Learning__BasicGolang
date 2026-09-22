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
// Объяви псевдоним type Code = int.
// Создай переменную типа Code со значением 201 и верни её.
func task1() int {
	return 0
}

// Задание 2
// Объяви новый тип type Port int — без знака =.
// Создай значение Port со числом 443.
// Верни его, преобразовав к int.
func task2() int {
	return 0
}

// Задание 3
// Объяви псевдоним type Bookmarks = map[string]string.
// Положи в такую map пару "docs" → "https://go.dev".
// Верни адрес по ключу "docs".
func task3() string {
	return ""
}
