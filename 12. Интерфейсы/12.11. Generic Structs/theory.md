# 12.11. Generic Structs

Generic struct — struct с параметром типа. Поле хранит значение этого типа. Один `Box` подходит и для строки, и для числа. Конкретный тип пишут при создании значения.

## Поле типа T

```go
type Box[T any] struct {
	Value T
}
```

`[T any]` стоит у имени struct. `Value T` — поле того типа, который подставят снаружи. `any` снова значит: ограничение широкое, подойдёт любой тип.

```go
text := Box[string]{Value: "vault"}
n := Box[int]{Value: 4}
```

`Box[string]` — коробка со строкой. `Box[int]` — коробка с числом. Квадратные скобки здесь обязательны: компилятор должен знать, чем заменить `T`.

## Метод

Получатель метода повторяет параметр типа.

```go
func (b Box[T]) Get() T {
	return b.Value
}
```

`Box[T]` — получатель. Писать просто `Box` нельзя: у типа есть параметр. `Get` возвращает поле. `text.Get()` вернёт `"vault"`. `n.Get()` вернёт `4`.

## Более узкое поле

Если метод складывает числа, ограничение `any` уже не подходит. Для счётчика оставляют `int`.

```go
type Number interface {
	int
}

type Counter[T Number] struct {
	N T
}

func (c Counter[T]) Next() T {
	return c.N + 1
}
```

`Number` разрешает только `int`. Поэтому `c.N + 1` законно. `Counter[int]{N: 6}` и вызов `Next` дают `7`.

Сравнение с пустым значением требует `comparable`:

```go
type Slot[T comparable] struct {
	Value T
}

func (s Slot[T]) Filled() bool {
	var zero T
	return s.Value != zero
}
```

`var zero T` — нулевое значение типа `T`. Для строки это `""`, для числа это `0`. `!=` разрешён, потому что ограничение — `comparable`. `Slot[string]{Value: "x"}` и `Filled` дают `true`.

## Что запомнить

- У generic struct параметр типа стоит в имени: `type Box[T any] struct`.
- Значение создают с конкретным типом: `Box[string]{Value: "vault"}`.
- Получатель метода пишет `Box[T]`, не голое `Box`.
- Сложение и сравнение требуют узкое ограничение: `int` или `comparable`.

## Что сделать

Заполни `tasks.go`. Типы опиши на уровне пакета. Запуск: `go run .`. Проверка: `go test .`.
