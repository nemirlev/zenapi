# ZenMoney API SDK на Go.

[![GoDoc](https://godoc.org/github.com/zenapi/zenapi?status.svg)](https://godoc.org/github.com/nemirlev/zenapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/nemirlev/zenapi)](https://goreportcard.com/report/github.com/nemirlev/zenapi)
![GitHub License](https://img.shields.io/github/license/nemirlev/zenapi)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/nemirlev/zenapi)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/nemirlev/zenapi)

Это SDK предназначено для взаимодействия с ZenMoney API. На данный момент поддерживает только метод diff, так как
suggestion работает достаточно странно и не понятно на сколько он нужен.

## Установка

Чтобы установить этот пакет, вы можете использовать команду `go get`:

```bash
go get github.com/nemirlev/zenapi
```

## Использование

Получите токен через [Zerro.app](https://zerro.app/token). Передайте через переменные окружения `ZENMONEY_TOKEN`

Для использования этого SDK вам нужно импортировать его в ваш проект:

```go
import "github.com/nemirlev/zenapi"
```

Затем вы можете создать новый клиент и использовать его для вызова методов API:

```go
client, err := zenapi.NewClient()
if err != nil {
// обработка ошибки
}
```

## Методы

В настоящее время SDK поддерживает следующие методы:

* Sync(body Request) - синхронизация данных. Данный запрос используется для получения и отправки изменений в данных с
  момента
  последней синхронизации.
* FullSync() - полная синхронизация данных. Данный запрос используется для получения всех данных из ZenMoney. Возвращает
  структуру Response

## Лицензия

Этот проект лицензирован под лицензией MIT - подробности см. в файле LICENSE.

## Вклад в проект

Мы приветствуем вклад от сообщества! Если вы хотите внести изменения в код, пожалуйста, следуйте этим шагам:

1. Форкните репозиторий.
2. Создайте новую ветку для ваших изменений.
3. Сделайте изменения в вашей ветке.
4. Отправьте Pull Request с описанием ваших изменений.

Пожалуйста, убедитесь, что ваш код соответствует стандартам Go и что все тесты проходят перед отправкой PR.

> Если вы хотите помочь, но не знаете с чего начать, то посмотрите Issues и создайте свой, если не нашли подходящего.