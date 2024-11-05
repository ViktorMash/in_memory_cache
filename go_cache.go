package go_cache

// Cache - структура, содержащая map для хранения данных
type Cache struct {
	data map[string]interface{}
}

// Функция-конструктор NewCache создает и возвращает указатель на новый экземпляр Cache с инициализированной map.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.data[key]
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
