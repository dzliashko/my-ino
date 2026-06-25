# Задание №1

### Напиши функцию:

```
func isPopular(subscribers int) bool
```

### Правила:

- меньше 1000 → false
- 1000 и больше → true

### Проверь на значениях:

```
100
1000
2500
```

# Задание №2

### Напиши функцию:

```
func channelRating(subscribers int) string
```

### Правила:

```
< 100      -> New
100-999    -> Growing
1000-9999  -> Popular
10000+     -> Top
```

### Пример:

```
fmt.Println(channelRating(50))
fmt.Println(channelRating(500))
fmt.Println(channelRating(5000))
fmt.Println(channelRating(50000))
```

### Ожидается:

```
New
Growing
Popular
Top
```

# Задание №3 (обязательное)

### Напиши функцию:

```
func averageArticles(totalArticles int, days int) float64
```

### Пример:

```
averageArticles(60, 30)
```

### должно вернуть:

```
2.0
```

### Здесь ты впервые столкнёшься с преобразованием типов:

```
float64(totalArticles)
```
