# 15.8. Post запрос

`GET` читает данные. `POST` отправляет данные на сервер в теле запроса. Тело — снова `io.Reader`, его можно собрать из строки.

`http.Post` делает такой запрос.

```go
payload := strings.NewReader(`{"city":"Moscow"}`)
resp, err := http.Post(srv.URL, "application/json", payload)
if err != nil {
	return ""
}
defer resp.Body.Close()
```

По строкам:

- `strings.NewReader` — поток из текста JSON. JSON здесь — строка с полем `city`.
- Первый аргумент `http.Post` — адрес.
- Второй — заголовок `Content-Type`. `application/json` говорит, что тело является JSON.
- Третий — тело запроса, `io.Reader`.
- Ответ устроен так же, как у `GET`: ошибка, затем `defer resp.Body.Close()`.

## Что видит сервер

Учебный сервер читает пришедший запрос и пишет ответ.

```go
srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method))
}))
defer srv.Close()
```

`r.Method` — метод запроса строкой. Для `http.Post` это `POST`. `w.Write` кладёт эту строку в тело ответа. Клиент читает её через `io.ReadAll`.

Заголовок тела лежит в запросе:

```go
kind := r.Header.Get("Content-Type")
```

`Header` — набор заголовков. `Get` возвращает значение по имени. После `http.Post` с вторым аргументом `application/json` здесь будет `application/json`.

Само тело запроса тоже `Reader`:

```go
data, err := io.ReadAll(r.Body)
```

`r.Body` — то, что клиент отправил. `string(data)` для текста `{"city":"Moscow"}` даст ту же строку.

## Что запомнить

- `POST` отправляет тело. `http.Post` принимает адрес, `Content-Type` и `io.Reader`.
- JSON-тело для этого клиента: `{"city":"Moscow"}`. Тип содержимого: `application/json`.
- Ответ закрывают через `defer resp.Body.Close()`.
- На сервере метод лежит в `r.Method`, заголовок — в `r.Header.Get`, тело — в `r.Body`.

## Что сделать

Заполни `tasks.go`. В каждом задании свой `httptest`-сервер и один `http.Post`. Запуск: `go run .`. Проверка: `go test .`.
