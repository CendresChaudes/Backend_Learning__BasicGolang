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
// В ответ запиши r.Method.
// Отправь http.Post на srv.URL.
// Content-Type: "application/json".
// Тело запроса: пустой Reader, strings.NewReader("").
// Прочитай тело ответа через io.ReadAll.
// Если запрос или чтение вернули ошибку, верни пустую строку.
// Иначе верни тело ответа.
// Сервер и тело ответа закрой через defer.
func task1() string {
	return ""
}

// Задание 2
// Подними сервер через httptest.NewServer.
// В ответ запиши r.Header.Get("Content-Type").
// Отправь http.Post на srv.URL с Content-Type "application/json".
// Тело: strings.NewReader("").
// Прочитай тело ответа.
// Если запрос или чтение вернули ошибку, верни пустую строку.
// Иначе верни тело ответа.
// Сервер и тело ответа закрой через defer.
func task2() string {
	return ""
}

// Задание 3
// Подними сервер через httptest.NewServer.
// Прочитай r.Body через io.ReadAll и запиши эти байты в ответ.
// Если чтение тела запроса вернуло ошибку, напиши в ответ пустой срез.
// Отправь http.Post на srv.URL с Content-Type "application/json".
// Тело запроса — ровно {"city":"Moscow"}, без перевода строки.
// Прочитай тело ответа.
// Если запрос или чтение ответа вернули ошибку, верни пустую строку.
// Иначе верни тело ответа.
// Сервер и тело ответа закрой через defer.
func task3() string {
	return ""
}
