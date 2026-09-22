package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Опиши метод Label() string у User.
// Он возвращает "user: " и сразу имя.
// Верни подпись пользователя "Ann".
func task1() string {
	return ""
}

// Задание 2
// Опиши метод Rename(name string) с получателем *User.
// Он записывает новое имя в исходный struct.
// Создай User с именем "Ann", переименуй в "Mia" и верни имя.
func task2() string {
	return ""
}

// Задание 3
// Опиши метод Older() с получателем *User.
// Он увеличивает возраст на 1.
// Создай User с возрастом 20, вызови Older и верни возраст.
func task3() int {
	return 0
}
