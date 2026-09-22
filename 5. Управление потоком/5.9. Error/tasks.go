package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(""))

	fmt.Println("Задание 2")
	fmt.Println(task2(10, 0))

	fmt.Println("Задание 3")
	fmt.Println(task3(20))
}

// Задание 1
// Если name пустое, создай ошибку errors.New("пустое имя") и верни её текст.
// Если имя не пустое, верни пустую строку: ошибки нет.
func task1(name string) string {
	return ""
}

// Задание 2
// Раздели a на b.
// Если b равен 0, верни текст ошибки "деление на ноль".
// Иначе верни частное через fmt.Sprintf("%d", ...).
// Ошибку создай через errors.New.
func task2(a, b int) string {
	return ""
}

// Задание 3
// Код ответа бывает от 100 до 599 включительно.
// Если code вне этого диапазона, верни текст ошибки "плохой код".
// Иначе верни "ok". Ошибку создай через errors.New.
func task3(code int) string {
	return ""
}
