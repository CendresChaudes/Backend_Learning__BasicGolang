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
// Подними сервер через httptest.NewServer.
// В тело ответа запиши "ясно".
// Сделай http.Get на srv.URL и прочитай тело через io.ReadAll.
// Если запрос или чтение вернули ошибку, верни пустую строку.
// Иначе верни тело ответа строкой.
// Сервер и тело ответа закрой через defer.
func task1() string {
	return ""
}

// Задание 2
// Подними сервер через httptest.NewServer.
// Поставь статус 404 через WriteHeader.
// Сделай http.Get на srv.URL.
// Если запрос вернул ошибку, верни 0.
// Иначе верни resp.StatusCode.
// Сервер и тело ответа закрой через defer.
func task2() int {
	return 0
}

// Задание 3
// Подними сервер через httptest.NewServer.
// Статус не ставь: без WriteHeader он будет 200.
// Сделай http.Get на srv.URL.
// Верни true, если resp.StatusCode равен 200.
// Если запрос вернул ошибку, верни false.
// Сервер и тело ответа закрой через defer.
func task3() bool {
	return false
}
