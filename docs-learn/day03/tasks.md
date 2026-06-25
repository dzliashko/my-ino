# Задание №1

### Создай структуру:

```
type Feed struct {
Name string
URL string
Subscribers int
Active bool
}
```

### Задание №2

В main() создай два канала:

```
Go Blog
https://go.dev/blog/feed.atom
1500 подписчиков
активный
```

и

```
Hacker News
https://hnrss.org/frontpage
50000 подписчиков
активный
```

### Задание №3

Выведи информацию о каждом канале.

Пример:

```
Feed: Go Blog
Subscribers: 1500
```

### Задание №4 (важное)

Используй функцию channelRating() из прошлого урока.

То есть:

```
fmt.Println(channelRating(feed.Subscribers))
```

Для каждого канала вывести рейтинг.

Пример:

```
Go Blog -> Popular
Hacker News -> Top
```

### Дополнительное задание

Добавь поле:

```
Articles int
```

в структуру Feed.

И выведи количество статей.
