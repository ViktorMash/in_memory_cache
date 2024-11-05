# in_memory_cache Package

`in_memory_cache` — это простой пакет для создания и управления кэшем на основе `map` в Go. Пакет позволяет добавлять, извлекать и удалять значения по ключу.

## Установка

Для установки этого пакета выполните команду:

```bash
go get github.com/ViktorMash/in_memory_cache


## Описание

Пакет `in_memory_cache` предоставляет следующие функции:

- `NewCache() *Cache` — функция-конструктор, которая создает новый экземпляр `Cache` с инициализированной `map`.
- `Set(key string, value interface{})` — метод для добавления или обновления значения в кэше.
- `Get(key string) interface{}` — метод для получения значения по ключу из кэша.
- `Delete(key string)` — метод для удаления значения из кэша по ключу.

## Пример использования

### Импорт пакета

```go
package main

import (
	"fmt"

	"github.com/ViktorMash/in_memory_cache"
)

func main() {
	// Создаем новый экземпляр Cache
	cache := in_memory_cache.NewCache()

	// Добавляем значения в кэш
	cache.Set("name", "Alice")
	cache.Set("age", 30)

	// Извлекаем значение по ключу
	name := cache.Get("name")
	fmt.Println("Name:", name)

	age := cache.Get("age")
	fmt.Println("Age:", age)

	// Обновляем значение
	cache.Set("age", 31)
	new_age := cache.Get("age")
	fmt.Println("Age:", new_age) // Output: Age: 31

	// Удаляем значение
	cache.Delete("name")
	name = cache.Get("name")
	if name == nil {
		fmt.Println("Name not found")
	}
}
```