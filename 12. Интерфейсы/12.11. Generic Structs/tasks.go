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
// Опиши type Box[T any] struct с полем Value T.
// Метод Get с получателем Box[T] возвращает Value.
// Создай Box[string] со значением "vault" и верни Get.
func task1() string {
	return ""
}

// Задание 2
// Опиши интерфейс Number с типом int.
// Опиши type Counter[T Number] struct с полем N T.
// Метод Next с получателем Counter[T] возвращает N + 1.
// Создай Counter[int] с N равным 6 и верни Next.
func task2() int {
	return 0
}

// Задание 3
// Опиши type Slot[T comparable] struct с полем Value T.
// Метод Filled возвращает true, если Value не равно нулевому T.
// Нулевое значение получи через var zero T.
// Создай Slot[string] со значением "x" и верни Filled.
func task3() bool {
	return false
}
