# 15.6. Query параметры

Город не зашивают в путь навсегда. Его передают query-параметром. Query — часть URL после знака `?`.

```text
https://api.weather.local/v1?city=Moscow&units=C
```

`city=Moscow` — пара «имя=значение». `&` отделяет следующую пару `units=C`.

## Набор пар

Пакет `net/url` собирает такие пары. `url.Values` — набор: у каждого имени есть значение.

```go
q := url.Values{}
q.Set("city", "Moscow")
q.Set("units", "C")
query := q.Encode()
```

По строкам:

- `url.Values{}` — пустой набор.
- `Set` записывает одну пару. Повторный `Set` с тем же именем заменяет значение.
- `Encode` делает одну строку. Имена идут по алфавиту. Здесь получится `city=Moscow&units=C`.

Пробел в значении `Encode` записывает как `+`. Город `New York` станет `city=New+York`.

## Адрес целиком

Базовый адрес и query склеивают через `?`.

```go
rawURL := "https://api.weather.local/v1?" + query
```

В `rawURL` будет `https://api.weather.local/v1?city=Moscow&units=C`. Эту строку потом передают в `http.Get`.

## Разобрать обратно

Чужую query-строку читают через `url.ParseQuery`.

```go
values, err := url.ParseQuery("city=Kazan&units=F")
if err != nil {
	return ""
}
city := values.Get("city")
```

`ParseQuery` возвращает `url.Values` и ошибку. `Get` отдаёт значение по имени. Для ключа `city` это `Kazan`. Если имени нет, `Get` вернёт пустую строку.

## Что запомнить

- Query стоит в URL после `?`. Пары разделяет `&`.
- `url.Values` хранит пары. `Set` записывает, `Encode` собирает строку, ключи по алфавиту.
- Полный адрес: базовый URL, `?`, результат `Encode`.
- `url.ParseQuery` разбирает строку обратно. `Get` читает одно значение.

## Что сделать

Заполни `tasks.go`. Запуск: `go run .`. Проверка: `go test .`.
