# 12.5. Встроенный интерфейс

Встроенный интерфейс уже описан в Go. Свой список методов для него писать не нужно. Достаточно сделать у своего типа методы с теми же именами. Тогда значение этого типа принимают функции стандартной библиотеки.

## error

`error` — интерфейс с одним методом:

```go
type error interface {
	Error() string
}
```

Любой тип с методом `Error() string` подходит под `error`. Текст ошибки возвращает этот метод.

```go
type Missing struct {
	Path string
}

func (m Missing) Error() string {
	return "нет " + m.Path
}
```

`Missing` хранит путь. `Error` склеивает `"нет "` и путь. Для `"notes.txt"` получится `"нет notes.txt"`.

Функция принимает уже готовый интерфейс `error`:

```go
func text(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
```

`err == nil` значит, что ошибки нет. Метод тогда не вызывают: у пустого `error` нет значения, вызов `Error` остановит программу. Если ошибка есть, `err.Error()` отдаёт текст.

```go
msg := text(Missing{Path: "notes.txt"})
```

В `msg` будет `"нет notes.txt"`. `Missing` передан туда, где ждут `error`, потому что метод совпал.

## fmt.Stringer

`fmt.Stringer` — интерфейс пакета `fmt`. В нём один метод: `String() string`.

```go
type Entry struct {
	Name string
}

func (e Entry) String() string {
	return "entry " + e.Name
}
```

`String` возвращает короткую подпись. Для имени `"mail"` это `"entry mail"`.

`fmt.Println` смотрит, есть ли у значения метод `String`. Если есть, в вывод идёт его результат. Вызов метода напрямую выглядит так: `e.String()`.

## Что запомнить

- `error` — готовый интерфейс. Нужен метод `Error() string`.
- Свой тип с таким методом передают туда, где аргумент имеет тип `error`.
- Пока `err` равен `nil`, метод `Error` не вызывают.
- `fmt.Stringer` — готовый интерфейс с методом `String() string`. Его использует печать из `fmt`.

## Что сделать

Заполни `tasks.go`. Типы `Missing` и `Entry` опиши на уровне пакета. Запуск: `go run .`. Проверка: `go test .`.
