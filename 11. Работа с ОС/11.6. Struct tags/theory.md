# 11.6. Struct tags

Имя поля в Go и имя поля в JSON часто различаются. В Go принято `Login` с большой буквы. В JSON для такого значения пишут `login`. Тег struct говорит пакету `json`, какое имя использовать.

## Тег

Тег — строка в обратных кавычках сразу после типа поля.

```go
type Account struct {
	Login  string `json:"login"`
	Secret string `json:"secret"`
}

data, err := json.Marshal(Account{Login: "ann", Secret: "qwe"})
```

`` `json:"login"` `` читает пакет `encoding/json`. Внутри кавычек — имя поля в JSON. При `json.Marshal` поле `Login` станет `"login"`, поле `Secret` станет `"secret"`.

Если `err` равен `nil`, в `data` лежит текст `{"login":"ann","secret":"qwe"}`.

## Пропуск поля

Тег `json:"-"` исключает поле из JSON. Минус — одно слово внутри кавычек.

```go
type Account struct {
	Login string `json:"login"`
	Code  string `json:"-"`
}
```

`Code` в текст не попадёт, даже если в нём лежит строка. `json.Marshal` значения с логином `"ann"` и любым кодом даст `{"login":"ann"}`.

## Пустое поле

`,omitempty` дописывают к имени через запятую. Пустое значение в JSON не попадает. Для строки пустое — `""`. Для числа пустое — `0`.

```go
type Account struct {
	Login    string `json:"login"`
	Nickname string `json:"nickname,omitempty"`
}
```

Если `Nickname` равен `""`, в тексте останется только логин: `{"login":"ann"}`. Если ник задан, оба поля будут в JSON.

## Что запомнить

- Тег `` `json:"имя"` `` задаёт имя поля в JSON.
- `` `json:"-"` `` убирает поле из JSON.
- `` `json:"имя,omitempty"` `` пропускает пустое значение.
- Без тега в JSON уходит имя поля Go.

## Что сделать

Заполни `tasks.go`. Запуск: `go run .`. Проверка: `go test .`.
