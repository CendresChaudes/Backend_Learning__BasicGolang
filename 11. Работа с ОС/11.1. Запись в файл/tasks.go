package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("note.txt", "go"))

	fmt.Println("Задание 2")
	fmt.Println(task2("perm.txt"))

	fmt.Println("Задание 3")
	fmt.Println(task3("note.txt", "go"))
}

// Задание 1
// Запиши text в файл path через os.WriteFile и права 0644.
// Верни число байтов, которые попали в файл.
// Если text пустой, файл не создавай и верни 0.
func task1(path, text string) int {
	return 0
}

// Задание 2
// Запиши в path байты строки "ok" с правами 0644.
// Верни эти права числом: 0644.
func task2(path string) int {
	return 0
}

// Задание 3
// Если text пустой, верни "пустой текст" и ничего не записывай.
// Иначе запиши text в path и верни пустую строку.
func task3(path, text string) string {
	return ""
}
