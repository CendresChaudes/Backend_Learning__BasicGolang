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
// Создай пакет files, импорт "lesson10/bundle/files".
// В файле file.go опиши struct File с полями Name string и Size int.
// Оба поля экспортированы.
// Создай File с именем "notes.txt" и размером 12. Верни имя.
func task1() string {
	return ""
}

// Задание 2
// В файле size.go того же пакета напиши функцию Total(a, b File) int.
// Она возвращает сумму полей Size.
// Возьми два файла с размерами 12 и 8. Верни files.Total.
func task2() int {
	return 0
}

// Задание 3
// В файле label.go того же пакета напиши функцию Label(f File) string.
// Подключи fmt в этом файле.
// Label возвращает fmt.Sprintf("%s %d", f.Name, f.Size).
// Верни подпись файла "notes.txt" размера 12.
func task3() string {
	return ""
}
