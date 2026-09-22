# 11.2. JSON

JSON — текстовый формат для данных. Его удобно хранить в файле и читать из другой программы. Значение Go превращается в такой текст и собирается обратно.

Объект JSON записан в фигурных скобках. Внутри пары: имя поля в кавычках, двоеточие, значение.

```text
{"login":"ann","secret":"qwe"}
```

Массив — несколько значений в квадратных скобках.

## Marshal и теги

Пакет `encoding/json` переводит значение в байты. `json.Marshal` возвращает байты и ошибку.

Имя поля в Go и имя поля в JSON часто различаются. В Go принято `Login` с большой буквы. В JSON пишут `login`. Тег struct говорит пакету `json`, какое имя использовать. Тег — строка в обратных кавычках сразу после типа поля.

```go
type Account struct {
	Login  string `json:"login"`
	Secret string `json:"secret"`
}

data, err := json.Marshal(Account{Login: "ann", Secret: "qwe"})
```

Если `err` равен `nil`, в `data` лежит текст `{"login":"ann","secret":"qwe"}`.

Тег `` `json:"-"` `` убирает поле из JSON. `` `json:"имя,omitempty"` `` пропускает пустое значение: для строки это `""`, для числа — `0`.

Без тега в JSON уходит имя поля Go.

## Файл и срез

Байты из `Marshal` отдают `os.WriteFile`. Обратно `os.ReadFile` читает байты, а `json.Unmarshal` кладёт их в переменную. Адрес переменной передают через `&`.

Несколько записей — срез struct. Тип среза пишут `[]Entry`.

```go
type Entry struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

var items []Entry
err := json.Unmarshal(data, &items)
```

Текст `[{"name":"mail","secret":"a"}]` становится срезом из одной записи. Поле `items[0].Name` равно `"mail"`.

## Что запомнить

- `json.Marshal` делает байты JSON. `json.Unmarshal` собирает значение обратно.
- Тег `` `json:"имя"` `` задаёт имя поля в JSON.
- Срез struct соответствует массиву объектов JSON.

## Что сделать

Заполни `tasks.go`. Запуск: `go run .`. Проверка: `go test .`.
