package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("note.txt", "привет"))

	fmt.Println("Задание 2")
	fmt.Println(task2("missing.txt"))

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Запиши text в файл path через os.WriteFile. Права 0644.
// Затем прочитай этот файл и верни текст.
// Если запись или чтение вернули ошибку, верни пустую строку.
func task1(path, text string) string {
	return ""
}

// Задание 2
// Прочитай файл path.
// Если ошибки нет, верни текст файла.
// Если ошибка есть, верни "нет файла".
func task2(path string) string {
	return ""
}

// Задание 3
// Верни строку "открыл|закрыл".
// Результат функции сделай именованным.
// Сначала запиши "открыл". Через defer допиши "|закрыл".
func task3() string {
	return ""
}
