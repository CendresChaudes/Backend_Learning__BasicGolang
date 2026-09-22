# 14.6. Упражнение Применение шифрования

Запись хранилища по-прежнему состоит из имени и секрета.

```go
type Entry struct {
	Name   string
	Secret string
}
```

В файл должен уйти не `Secret`, а его шифротекст. Ключ берут из переменной окружения `VAULT_KEY`, не из кода.

```bash
VAULT_KEY=1234567890123456 go run .
```

`os.Getenv("VAULT_KEY")` возвращает этот ключ строкой. Из строки собирают `Encrypter`.

## Спрятать секрет

`Encrypt` тот же, что в уроке про шифрование. Учебный nonce снова `[]byte("0123456789ab")`.

```go
func (e Encrypter) Encrypt(plain string) (string, error) {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := []byte("0123456789ab")
	sealed := gcm.Seal(nil, nonce, []byte(plain), nil)
	full := append(nonce, sealed...)
	return hex.EncodeToString(full), nil
}
```

Для записи вызов выглядит так:

```go
key := os.Getenv("VAULT_KEY")
enc := NewEncrypter(key)
hidden, err := enc.Encrypt(entry.Secret)
```

`entry.Secret` — открытый секрет, например `"b2"`. В `hidden` попадает hex-строка. Её и сохраняют вместо пароля. `entry.Name` шифровать не нужно: по имени запись потом ищут.

## Открыть секрет

`Decrypt` забирает hex-строку и тот же ключ из окружения.

```go
func (e Encrypter) Decrypt(encoded string) (string, error) {
	data, err := hex.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	if len(data) < 12 {
		return "", errors.New("короткие данные")
	}
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := data[:12]
	sealed := data[12:]
	plain, err := gcm.Open(nil, nonce, sealed, nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}
```

Проверка, что секрет не испортился: зашифровать и сразу расшифровать. Если открытый текст снова равен `entry.Secret`, имя записи можно показать пользователю.

```go
hidden, err := enc.Encrypt(entry.Secret)
if err != nil {
	return ""
}
opened, err := enc.Decrypt(hidden)
if err != nil || opened != entry.Secret {
	return ""
}
return entry.Name
```

Для `Entry{Name: "mail", Secret: "anna"}` такой проход вернёт `"mail"`. Пустая строка значит: ключа не хватило или шифр не сошёлся.

## Что запомнить

- Ключ читают из `VAULT_KEY` через `os.Getenv`.
- В хранилище записывают hex из `Encrypt`, не открытый `Secret`.
- `Decrypt` тем же ключом возвращает секрет.
- Имя записи не шифруют. Его возвращают, когда секрет после круга снова совпал.

## Что сделать

Это упражнение. Тип `Entry` уже описан в `tasks.go`. Методы шифрования опиши на уровне пакета. Запуск: `VAULT_KEY=1234567890123456 go run .`. Проверка: `go test .`.
