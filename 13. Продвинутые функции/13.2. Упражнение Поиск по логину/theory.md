# 13.2. Упражнение Поиск по логину

В хранилище одна запись — логин и секрет.

```go
type Account struct {
	Login  string
	Secret string
}
```

Поиск не должен знать, по какому полю сравнивать. Правило сравнения передают функцией. Тип такого аргумента: `func(Account, string) bool`. Функция получает запись и строку запроса, возвращает `bool`.

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

В вызове третьим аргументом передают анонимную функцию: она возвращает `true`, когда `item.Login` равен `login`. Та же схема находит сам факт «запись есть» и индекс записи.

## Что запомнить

- Правило «эта запись подходит» — функция с типом `func(Account, string) bool`.
- Поиск вызывает её для каждой записи.
- Нет подходящей записи — пустая строка, `false` или `-1`.

## Что сделать

Заполни `tasks.go`. Запуск: `go run .`. Проверка: `go test .`.
