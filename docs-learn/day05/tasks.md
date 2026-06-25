# День 5. Слайсы (slice)

До этого у нас был один канал:
```
goFeed
hnFeed
```
Но Inoreader работает с сотнями каналов.

Нужен список.

В Go для этого используется slice.
```
feeds := []Feed{
	goFeed,
	hnFeed,
}
```

### Теория

Создание слайса:
```
numbers := []int{1, 2, 3}
```
Обход:

```
for _, n := range numbers {
	fmt.Println(n)
}
```

range — один из самых используемых операторов в Go.

### Задание №1

Создай слайс:
```
feeds := []Feed{...}
```
из своих двух каналов.

### Задание №2

Выведи информацию о всех каналах через цикл:
```
for _, feed := range feeds {
	fmt.Print(feed.Info())
}
```
### Задание №3

Посчитай общее количество подписчиков.

Ожидаемая логика:
```
totalSubscribers := 0

for _, feed := range feeds {
	totalSubscribers += feed.Subscribers
}
```
### Задание №4

Напиши функцию:
```
func totalSubscribers(feeds []Feed) int
```
которая считает подписчиков всех каналов.

### Задание №5 (самое важное)

Напиши функцию:
```
func popularFeeds(feeds []Feed) []Feed
```
которая возвращает только популярные каналы.

Подсказка:
```
var result []Feed

for _, feed := range feeds {
	if feed.IsPopular() {
		result = append(result, feed)
	}
}
```
Здесь ты впервые познакомишься с:
```
append()
```
Это одна из самых важных функций в Go.

После этого мы сможем перейти к следующей сущности проекта — Article, а там появятся даты (time.Time) и вложенные структуры.
