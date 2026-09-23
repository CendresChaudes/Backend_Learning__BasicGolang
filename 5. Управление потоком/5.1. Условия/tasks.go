package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(200))

	fmt.Println("Задание 2")
	fmt.Println(task2(20, true))

	fmt.Println("Задание 3")
	fmt.Println(task3(75))
}

// Задание 1
// Если code равен 200, верни "ok". Иначе верни "ошибка".
func task1(code int) string {
	if code == 200 {
		return "ok"
	}

	return "ошибка"
}

// Задание 2
// Верни true, только если age больше или равен 18 и active равен true.
func task2(age int, active bool) bool {
	return age >= 18 && active
}

// Задание 3
// Если score больше или равен 90, верни "отлично".
// Иначе если score больше или равен 70, верни "хорошо".
// Иначе верни "ещё раз".
func task3(score int) string {
	var result string
	if score >= 90 {
		result = "отлично"
	} else if score >= 70 {
		result = "хорошо"
	} else {
		result = "ещё раз"
	}

	return result
}
