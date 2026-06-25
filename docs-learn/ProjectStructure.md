Мы не будем создавать "корпоративный монстр" из 20 папок, а сделаем структуру, которая соответствует текущему состоянию проекта.

# Цель

Из этого:
```
main.go (250 строк)
```
получить:
```
rss-reader/
│
├── go.mod
│
├── cmd/
│   └── rss-reader/
│       └── main.go
│
└── internal/
    ├── domain/
    │   ├── feed.go
    │   └── article.go
    │
    └── service/
        ├── feed.go
        └── article.go
```
Обрати внимание — **никаких repository, handler, api, database** пока нет. Они появятся тогда, когда действительно понадобятся.

## Шаг 1. Создаём проект

Если ещё не сделал:
```
mkdir rss-reader
cd rss-reader

go mod init rss-reader
```
Теперь модуль называется
```
rss-reader
```
## Шаг 2. Создаём папки
```
rss-reader/
│
├── cmd/
│   └── rss-reader/
│
└── internal/
    ├── domain/
    └── service/
```
Почему cmd?

Потому что приложение запускается отсюда.

Позже у нас могут быть
```
cmd/rss-reader
cmd/importer
cmd/migrator
```
Каждая папка — отдельная исполняемая программа.

## Шаг 3. Переносим Feed

Создай
```
internal/domain/feed.go
```
```
package domain

type Feed struct {
	Name        string
	URL         string
	Subscribers int
	Active      bool
	Articles    int
}
```
Заметь:

**никаких функций здесь нет.**

## Шаг 4. Переносим Article
```
internal/domain/article.go
```
```
package domain

import "time"

type Article struct {
	Title       string
	Description string
	Link        string
	PublishedAt time.Time
	Feed        Feed
}
```
Теперь все модели лежат в одном месте.

## Шаг 5. Переносим бизнес-логику Feed

Создай
```
internal/service/feed.go
```
```
package service

import (
	"fmt"

	"rss-reader/internal/domain"
)
```
И перенеси сюда
```
func ChannelRating(...)
func (f domain.Feed) Info()
func (f domain.Feed) Rating()
func (f domain.Feed) IsPopular()
func TotalSubscribers(...)
func PopularFeeds(...)
```
### И вот здесь мы остановимся.

Потому что **код не скомпилируется.**

### И это специально.

Почему?

Попробуй сам ответить.

Например,
```
func (f domain.Feed) Info() string
```
не скомпилируется.

### Знаешь почему?

### Потому что...

Go запрещает определять методы для типа из другого пакета.

То есть нельзя сделать
```
package service

func (f domain.Feed) Info() {}
```
Это одно из самых важных правил языка.

## Значит, архитектура неправильная?

Да.

И это отличный урок.

Очень многие новички пытаются вынести методы в service, а потом удивляются ошибке компиляции.

## Правильная структура

Именно поэтому в Go методы модели лежат рядом с моделью.

Получается:
```
internal/
│
├── domain/
│   ├── feed.go
│   ├── feed_methods.go
│   ├── article.go
│   └── article_methods.go
│
└── service/
```
То есть:
```
package domain
```
в обоих файлах.

Например

**feed.go**
```
type Feed struct {
    ...
}
```
**feed_methods.go**
```
func (f Feed) Info() string
func (f Feed) Rating() string
func (f Feed) IsPopular() bool
```
Это уже настоящий Go-стиль.

## Тогда что лежит в service?

Не методы.

А бизнес-операции.

Например:
```
func PopularFeeds(...)
func TotalSubscribers(...)
func ArticlesByFeed(...)
```
Они работают сразу с несколькими объектами или коллекциями и не принадлежат какой-то одной сущности.

В итоге получится такая структура
```
rss-reader/
│
├── go.mod
│
├── cmd/
│   └── rss-reader/
│       └── main.go
│
└── internal/
    ├── domain/
    │   ├── feed.go
    │   ├── feed_methods.go
    │   ├── article.go
    │   └── article_methods.go
    │
    └── service/
        ├── feed.go
        └── article.go
```
