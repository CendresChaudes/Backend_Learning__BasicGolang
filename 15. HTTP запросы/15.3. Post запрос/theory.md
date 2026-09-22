# 15.3. Post запрос

`GET` читает данные. `POST` отправляет данные на сервер в теле запроса. Тело — снова `io.Reader`, его можно собрать из строки.

`http.Post` делает такой запрос. Второй аргумент — заголовок `Content-Type`, тип тела. Для JSON пишут `"application/json"`.

```go
payload := strings.NewReader(`{"city":"Moscow"}`)
resp, err := http.Post(srv.URL, "application/json", payload)
if err != nil {
	return ""
}
defer resp.Body.Close()
```

Сервер в учебном тесте поднимают через `httptest.NewServer`. Обработчик видит метод `r.Method`, заголовок `r.Header.Get("Content-Type")` и тело `r.Body`. Тело запроса тоже читают через `io.ReadAll`.

Ответ сервера читают так же, как у GET: `io.ReadAll(resp.Body)`. Сервер закрывают через `defer srv.Close()`.

## Что запомнить

- `POST` отправляет тело. `GET` тело обычно не несёт.
- Тело запроса — `io.Reader`. Для строки его делает `strings.NewReader`.
- `Content-Type` говорит серверу, как читать тело. Для JSON это `application/json`.

## Что сделать

Заполни `tasks.go`. Запуск: `go run .`. Проверка: `go test .`.
