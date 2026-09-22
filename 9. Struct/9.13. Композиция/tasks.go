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
// На уровне пакета опиши type Admin struct: встрой User и добавь поле Level int.
// Создай Admin с именем "Iva" и Level 3.
// Верни имя по короткому полю, без записи через User.
func task1() string {
	return ""
}

// Задание 2
// Используй те же User и Admin, что в задании 1.
// Верни Level этого администратора: имя "Iva", Level 3.
func task2() int {
	return 0
}

// Задание 3
// Опиши type Stats struct с полем Score int.
// Опиши type Player struct: встрой Stats и добавь поле Bonus int.
// Создай игрока: Score 10, Bonus 5.
// Верни сумму Score и Bonus. Score прочитай по короткому имени.
func task3() int {
	return 0
}
