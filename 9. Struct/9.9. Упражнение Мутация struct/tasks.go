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
// На уровне пакета опиши type Wallet struct с полем Balance int.
// Метод Deposit(amount int) с получателем *Wallet прибавляет amount к Balance.
// Начни с баланса 100 и вызови Deposit(25). Верни Balance.
func task1() int {
	return 0
}

// Задание 2
// На уровне пакета опиши type Profile struct с полем Name string.
// Метод Rename(name string) с получателем *Profile записывает новое имя.
// Начни с имени "Ann", вызови Rename("Nora"). Верни Name.
func task2() string {
	return ""
}

// Задание 3
// На уровне пакета опиши type Counter struct с полем N int.
// Метод Add(n int) с получателем *Counter прибавляет n к N.
// Начни с 10. Вызови Add(5), затем Add(7). Верни N.
func task3() int {
	return 0
}
