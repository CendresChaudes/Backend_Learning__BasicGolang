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
// Создай пакет user, импорт "lesson10/export/user".
// Экспортированная функция Login() string возвращает "Leo".
// Верни user.Login().
func task1() string {
	return ""
}

// Задание 2
// В пакете user опиши экспортированный struct Account
// с экспортированным полем Balance int.
// В task2 создай Account с Balance 80 и верни это поле.
func task2() int {
	return 0
}

// Задание 3
// В пакете user опиши struct Profile с полем name string.
// Поле name с маленькой буквы.
// Конструктор New(name string) Profile заполняет это поле.
// Метод Name() string возвращает его.
// Верни имя у user.New("Ann").
func task3() string {
	return ""
}
