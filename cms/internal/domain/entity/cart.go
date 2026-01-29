package entity

import "time"

type Cart struct {
	Cart_ID             uint64 `gorm:"primaryKey;autoIncrement"`
	quantity            int
	price               float64
	add_date            time.Time
	products_product_id uint64
	clients_client_id   uint64
}

func NewCart(quantity int, price float64, add_date time.Time, products_product_id uint64, clients_client_id uint64) *Cart {
	return &Cart{
		quantity:            quantity,
		price:               price,
		add_date:            add_date,
		products_product_id: products_product_id,
		clients_client_id:   clients_client_id,
	}
}

func (c *Cart) SetID(ID uint64) {
	c.Cart_ID = ID
}
