# 12.7. Type Switch

Type switch — `switch`, который смотрит на тип значения внутри `any`. Ветки — это типы: `string`, `int` и общая ветка `default`.

## Имя типа

```go
func kind(v any) string {
	switch v.(type) {
	case string:
		return "string"
	case int:
		return "int"
	default:
		return "other"
	}
}
```

`v.(type)` пишут только в таком `switch`. Скобки спрашивают тип того, что лежит в `any`.

`case string` срабатывает, когда внутри строка. `kind("vault")` вернёт `"string"`. `case int` срабатывает для целого числа. `kind(7)` вернёт `"int"`. `default` забирает всё остальное, например `bool`.

## Значение в ветке

Имя типа мало, если нужна длина строки или само число. Ветке отдают переменную нужного типа.

```go
func size(v any) int {
	switch x := v.(type) {
	case string:
		return len(x)
	case int:
		return x
	default:
		return 0
	}
}
```

`x := v.(type)` в каждой ветке имеет тип этой ветки. В `case string` переменная `x` — строка, у неё есть `len`. В `case int` переменная `x` — число, его возвращают как есть.

`size("vault")` вернёт `5`. `size(7)` вернёт `7`. Для другого типа функция вернёт `0`.

Проверка «это строка» строится так же:

```go
func isText(v any) bool {
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}
```

`isText("go")` вернёт `true`. `isText(7)` вернёт `false`.

## Что запомнить

- Type switch выбирает ветку по типу внутри `any`: `switch v.(type)`.
- `case` пишет тип: `string`, `int`. Остальное забирает `default`.
- `switch x := v.(type)` даёт в ветке значение уже этого типа.
- Форма `v.(type)` стоит только внутри такого `switch`.

## Что сделать

Заполни `tasks.go`. Запуск: `go run .`. Проверка: `go test .`.
