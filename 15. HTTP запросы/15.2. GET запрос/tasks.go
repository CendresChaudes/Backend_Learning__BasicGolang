package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("ясно"))

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3("Moscow", "C"))
}

// Задание 1
// Сделай Reader из text через strings.NewReader.
// Прочитай его через io.ReadAll.
// Если чтение вернуло ошибку, верни пустую строку.
// Иначе верни текст.
func task1(text string) string {
	return ""
}

// Задание 2
// Подними сервер через httptest.NewServer.
// В тело ответа запиши "ясно".
// Сделай http.Get на srv.URL и прочитай тело через io.ReadAll.
// Если запрос или чтение вернули ошибку, верни пустую строку.
// Иначе верни тело ответа строкой.
// Сервер и тело ответа закрой через defer.
func task2() string {
	return ""
}

// Задание 3
// Собери URL https://api.weather.local/v1 с query-параметрами city и units.
// Пары собери через url.Values и Encode.
// Верни полную строку адреса.
func task3(city, units string) string {
	return ""
}
