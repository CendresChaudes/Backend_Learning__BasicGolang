# 14.2. Шифрование

Секрет хранилища шифруют одним и тем же ключом: и когда прячут, и когда открывают. Ключ удобно держать в struct.

```go
type Encrypter struct {
	Key string
}

func NewEncrypter(key string) Encrypter {
	return Encrypter{Key: key}
}
```

AES — алгоритм шифрования из пакета `crypto/aes`. Ему подходит ключ длиной 16, 24 или 32 байта. Учебный ключ: `"1234567890123456"`.

GCM — режим из пакета `crypto/cipher`. Он шифрует байты и добавляет метку. Если шифротекст потом испортят, расшифровка это заметит.

## Nonce и Encrypt

Nonce — дополнительные 12 байт рядом с шифротекстом. Одним и тем же ключом нельзя шифровать два разных секрета с одним и тем же nonce. В настоящем хранилище nonce каждый раз новый, его берут из `crypto/rand` и сохраняют вместе с шифротекстом.

В уроке nonce фиксированный: `[]byte("0123456789ab")`. Так результат не меняется от запуска к запуску, и тест может сверить строку. В свой проект этот постоянный nonce не переносят.

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

`gcm.Seal` шифрует. Спереди к результату приписывают nonce. `encoding/hex` переводит байты в текст. Для секрета `"vault"` и этого ключа строка начинается с `303132333435363738396162`: это nonce, записанный через hex. Открытого слова `vault` в строке нет.

## Decrypt

Расшифровка режет hex обратно в байты. Первые 12 байт — nonce. Хвост отдаёт `gcm.Open`.

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
	plain, err := gcm.Open(nil, data[:12], data[12:], nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}
```

Ключ должен быть тем же. Чужая метка, чужой ключ или строка короче 12 байт — это ошибка, не секрет.

## Что запомнить

- Ключ живёт в `Encrypter`. Длина ключа AES — 16, 24 или 32 байта.
- `Encrypt` возвращает hex: сначала nonce, потом шифротекст.
- `Decrypt` тем же ключом возвращает исходную строку.
- Учебный nonce постоянный. В настоящем хранилище он каждый раз новый.

## Что сделать

Заполни `tasks.go`. Запуск: `go run .`. Проверка: `go test .`.
