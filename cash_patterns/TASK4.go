/*


Вы разрабатываете систему управления запасами для крупного интернет-магазина.
Ваша задача - оптимизировать процесс получения информации о товарах, используя стратегии кеширования.
Система должна уметь обрабатывать запросы на получение информации о товаре,
а также обновлять информацию о товаре в базе данных и кеше.

Требования:

Кеширование информации о товаре:
Когда пользователь запрашивает информацию о товаре,
система должна сначала проверять наличие этой информации в кеше.
Если информация отсутствует, она должна быть загружена из базы данных и сохранена в кеше.

Инвалидация кеша при изменении данных:
Когда информация о товаре обновляется (например, изменение цены или количества),
соответствующая запись в кеше должна быть инвалидирована или обновлена.

Применение стратегии TTL:
Для кешированных данных о товаре должен быть установлен TTL (Time-To-Live),
чтобы гарантировать актуальность информации.

Примечания

Структура Данных Товара:

type Product struct {
    ID    int
    Name  string
    Price float64
    Stock int
}

Функция получения информации о товаре:
func getProduct(productId int, db map[int]Product, cache *Cache) (Product, error)

Функция обновления информации о товаре (фейк-функция выполняющая роль базы данных):
func updateProduct(productId int, newProduct Product, db map[int]Product) error

Кеш продуктов:

type Cache struct {
    products map[int]Product  // Кэш продуктов
    ttl      time.Duration    // Время жизни записи в кэше
}

func NewCache() *Cache

Получение продукта из кеша:
func (c *Cache) Get(productId int) (Product, bool)

Сохранение продукта в кеш:
func (c *Cache) Set(productId int, product Product)

Инвалидация кеша:
func (c *Cache) Invalidate(productId int)

*/

package cash_patterns

import (
	"fmt"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func getProduct(productId int, db map[int]Product, cache *Cache) (Product, error) {
	val, ok := cache.Get(productId)
	if !ok {
		val, ok = db[productId]
		if !ok {
			return val, fmt.Errorf("Product doesnt exist")
		}
		cache.Set(productId, val)
	}
	return val, nil
}

func updateProduct(productId int, newProduct Product, db map[int]Product) error {
	_, ok := db[productId]
	if !ok {
		return fmt.Errorf("Product doesnt exist")
	}
	db[productId] = newProduct
	return nil
}

type Cache struct {
	products map[int]Product // Кэш продуктов
	ttl      time.Duration   // Время жизни записи в кэше
}

func NewCache() *Cache {
	return &Cache{
		products: make(map[int]Product),
		ttl:      time.Minute * 2,
	}
}

// Получение продукта из кеша:
func (c *Cache) Get(productId int) (Product, bool) {
	val, ok := c.products[productId]
	return val, ok
}

// Сохранение продукта в кеш:
func (c *Cache) Set(productId int, product Product) {
	c.products[productId] = product
}

// Инвалидация кеша:
func (c *Cache) Invalidate(productId int) {
	_, ok := c.products[productId]
	if ok {
		delete(c.products, productId)
	}
}
