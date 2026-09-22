# 11.10. Чтение JSON

Файл с паролями хранит JSON-массив: несколько объектов в квадратных скобках.

```text
[{"name":"mail","secret":"a"},{"name":"vpn","secret":"b"}]
```

Каждый объект — одна запись. Имена `name` и `secret` совпадают с тегами struct.

```go
type Entry struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}
```

## Два шага

Чтение собирает то, что ты уже делал по отдельности. Сначала байты файла, потом разбор.

```go
data, err := os.ReadFile(path)
if err != nil {
	if errors.Is(err, os.ErrNotExist) {
		return "нет файла"
	}
	return err.Error()
}

var list []Entry
err = json.Unmarshal(data, &list)
```

`var list []Entry` создаёт переменную среза. Пока в неё ничего не положили, длина равна нулю.

`&list` — адрес этой переменной. `json.Unmarshal` записывает в неё элементы массива. После успеха `len(list)` равен числу объектов в файле. У первой записи `list[0].Name` равен `"mail"`.

## Битый текст

Если байты — не JSON, `Unmarshal` вернёт ошибку. Её в этом уроке заменяют на короткую фразу:

```go
if err != nil {
	return "битый json"
}
```

Пока `err` равен `nil`, срез `list` уже заполнен. Дальше с ним работают как с обычным `[]Entry`: индекс, `len`, цикл.

Секрет нужной записи ищут по полю `Name`. Совпало имя — берут `Secret` этой записи.

## Что запомнить

- Файл паролей — JSON-массив объектов.
- `os.ReadFile` даёт байты, `json.Unmarshal` разбирает их в `[]Entry`.
- В `Unmarshal` передают адрес среза: `&list`.
- Нет файла — `"нет файла"`. Битый текст — `"битый json"`.

## Что сделать

Заполни `tasks.go`. Тип `Entry` уже описан. Запуск: `go run .`. Проверка: `go test .`.
