# 13.4. Упражнение. Поиск по логину

В хранилище одна запись — логин и секрет.

```go
type Account struct {
	Login  string
	Secret string
}
```

Раньше сравнение логина стояло прямо в цикле. Теперь сравнение — отдельная функция. Её передают в поиск. Цикл один и тот же, меняется только проверка.

Функция проверки принимает запись и логин. Логин приходит аргументом, рядом с записью.

```go
func(item Account, login string) bool {
	return item.Login == login
}
```

`item.Login == login` — логин записи совпал с искомым. Результат `bool`.

## Секрет

```go
func Secret(items []Account, login string, match func(Account, string) bool) string {
	for _, item := range items {
		if match(item, login) {
			return item.Secret
		}
	}
	return ""
}
```

`match` — переданная функция. Цикл отдаёт ей каждую запись и тот же `login`. Первый `true` останавливает поиск: наружу уходит `Secret` этой записи.

Вызов с анонимной функцией:

```go
secret := Secret(items, "boris", func(item Account, login string) bool {
	return item.Login == login
})
```

Если в списке есть `"boris"`, в `secret` будет его секрет. Если цикл закончился сам, логина нет. Тогда результат — пустая строка.

## Есть ли логин

Секрет для этой проверки не нужен. Достаточно `bool`.

```go
func Has(items []Account, login string, match func(Account, string) bool) bool {
	for _, item := range items {
		if match(item, login) {
			return true
		}
	}
	return false
}
```

Нашли запись — сразу `true`. Дошли до конца списка — `false`.

## Индекс

Индекс — позиция в срезе. Его даёт `i` в цикле `range`.

```go
func Index(items []Account, login string, match func(Account, string) bool) int {
	for i, item := range items {
		if match(item, login) {
			return i
		}
	}
	return -1
}
```

Первая запись имеет индекс `0`, поэтому ноль уже занят. Для «не найдено» возвращают `-1`: такого индекса в срезе не бывает.

## Что запомнить

- Проверка логина — функция `func(Account, string) bool`.
- В поиск её передают аргументом. Удобно передать анонимную функцию.
- Секрет первой подходящей записи. Нет записи — пустая строка.
- Индекс отсутствия — `-1`.

## Что сделать

Это упражнение. Тип `Account` уже описан в `tasks.go`. Функции `Secret`, `Has` и `Index` опиши на уровне пакета и вызови их из заданий. Запуск: `go run .`. Проверка: `go test .`.
