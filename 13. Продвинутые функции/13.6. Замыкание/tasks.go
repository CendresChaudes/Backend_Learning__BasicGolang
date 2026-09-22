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
// Опиши func withPrefix(prefix string) func(string) string.
// Она возвращает функцию. Та склеивает prefix и свой аргумент name.
// В task1 вызови withPrefix("vault:"), затем верни вызов результата с "mail".
func task1() string {
	return ""
}

// Задание 2
// Опиши func makeAdder(n int) func(int) int.
// Она возвращает функцию, которая прибавляет n к своему аргументу.
// В task2 вызови makeAdder(10), затем верни вызов результата с 5.
func task2() int {
	return 0
}

// Задание 3
// Опиши func counter() func() int.
// Внутри число n начинается с 0.
// Возвращённая функция прибавляет к n единицу и возвращает n.
// В task3 получи функцию из counter(). Вызови её один раз.
// Верни результат второго вызова.
func task3() int {
	return 0
}
