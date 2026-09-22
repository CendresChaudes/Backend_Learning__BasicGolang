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
// На уровне пакета опиши type User struct с полем Name string.
// Метод Label с получателем-значением User возвращает Name.
// Создай User с именем "Ann" и верни результат Label.
func task1() string {
	return ""
}

// Задание 2
// На уровне пакета опиши type Rect struct с полями W int и H int.
// Метод Area с получателем-значением возвращает произведение W и H.
// Создай прямоугольник 3 на 4 и верни Area.
func task2() int {
	return 0
}

// Задание 3
// На уровне пакета опиши type Guest struct с полем Age int.
// Метод IsAdult с получателем-значением возвращает true, если Age >= 18.
// Создай гостя с возрастом 20 и верни IsAdult.
func task3() bool {
	return false
}
