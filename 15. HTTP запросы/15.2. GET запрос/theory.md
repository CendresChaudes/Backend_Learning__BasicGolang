# 15.2. GET запрос

Ответ сервиса приходит потоком байтов. В Go такой поток читают через `io.Reader`. Reader — интерфейс с одним методом: из него можно читать байты.

Свой `Read` вызывать не нужно. `io.ReadAll` читает поток до конца и возвращает байты.

```go
src := strings.NewReader("ясно")
data, err := io.ReadAll(src)
text := string(data)
```

`strings.NewReader` делает `Reader` из готовой строки. Если ошибки нет, `text` равен `"ясно"`.

## Запрос

HTTP-запрос — сообщение клиента серверу. Метод `GET` значит «прочитать». Ответ — код статуса и тело. Тело как раз `Reader`.

Пакет `net/http` входит в стандартную библиотеку. `http.Get` отправляет `GET` по адресу.

```go
resp, err := http.Get(url)
if err != nil {
	return ""
}
defer resp.Body.Close()
data, err = io.ReadAll(resp.Body)
```

`resp.Body` — тело ответа, его тоже `Reader`. `defer` закрывает тело при выходе из функции. `resp.StatusCode` — число статуса. `200` значит успех.

В тесте настоящий сервис не нужен. `httptest.NewServer` поднимает локальный сервер на время вызова. В обработчике пишут ответ. `srv.URL` — его адрес. Сервер тоже закрывают через `defer srv.Close()`.

## Query

Город не зашивают в путь навсегда. Его передают query-параметром. Query — часть URL после знака `?`.

```text
https://api.weather.local/v1?city=Moscow&units=C
```

Пакет `net/url` собирает такие пары. `url.Values` — набор: у каждого имени есть значение.

```go
query := url.Values{}
query.Set("city", "Moscow")
query.Set("units", "C")
raw := "https://api.weather.local/v1?" + query.Encode()
```

`Encode` собирает строку пар. Имена он сортирует, поэтому `city` окажется раньше `units`.

## Что запомнить

- Тело ответа — `io.Reader`. Его читают через `io.ReadAll` и закрывают.
- `http.Get` отправляет GET. `StatusCode` — код ответа.
- Query собирают через `url.Values`, не склейкой строк вручную.

## Что сделать

Заполни `tasks.go`. Запуск: `go run .`. Проверка: `go test .`.
