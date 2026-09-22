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
// Создай рядом файл hello.go с первой строкой package main.
// В нём функция hello() string возвращает "привет".
// В task1 вызови hello и верни результат.
func task1() string {
	return ""
}

// Задание 2
// Создай пакет user, импорт "lesson10/packages/user".
// Экспортированная функция Login() string возвращает "Leo".
// Верни user.Login().
func task2() string {
	return ""
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
