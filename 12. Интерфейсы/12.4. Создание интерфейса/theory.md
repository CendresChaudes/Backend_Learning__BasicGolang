# 12.4. Создание интерфейса

Интерфейс — список методов. Тип подходит под интерфейс, когда у него есть все методы из этого списка. Совпадение имён, аргументов и результатов достаточно. Отдельная пометка на типе не нужна.

Так одна функция принимает и `Disk`, и `Cloud`. Ей важен метод, а не имя типа.

## Список методов

```go
type Labeler interface {
	Label() string
}
```

`type Labeler interface` задаёт имя интерфейса. Внутри одна строка: метод `Label` без аргументов, результат — `string`. Полей у интерфейса нет. Он описывает, что можно вызвать, а не из чего значение собрано.

## Функция принимает интерфейс

```go
func show(item Labeler) string {
	return item.Label()
}
```

Аргумент `item` имеет тип `Labeler`. Внутри `show` доступен вызов `item.Label()`. Поле `Name` отсюда не прочитать: в списке методов его нет.

Любое значение с методом `Label() string` можно передать в `show`.

```go
type Disk struct {
	Name string
}

func (d Disk) Label() string {
	return "disk:" + d.Name
}

type Cloud struct {
	Name string
}

func (c Cloud) Label() string {
	return "cloud:" + c.Name
}
```

У `Disk` и у `Cloud` есть нужный метод. Оба подходят под `Labeler`.

```go
disk := Disk{Name: "vault"}
cloud := Cloud{Name: "vault"}
fmt.Println(show(disk))
fmt.Println(show(cloud))
```

`show(disk)` вернёт `"disk:vault"`. `show(cloud)` вернёт `"cloud:vault"`. Тело `show` одно. Провайдер выбирают в аргументе.

## Что запомнить

- Интерфейс — список методов: `type Labeler interface { Label() string }`.
- Тип подходит, если такие методы у него уже есть.
- Функция с аргументом-интерфейсом вызывает только методы из списка.
- `Disk` и `Cloud` передают в одну функцию, если набор методов совпал.

## Что сделать

Заполни `tasks.go`. Интерфейс, типы и функцию `show` опиши на уровне пакета. Запуск: `go run .`. Проверка: `go test .`.
