package entity

type OrderedProducts struct {
	Quantity            uint8   `gorm:"column:quantity" json:"quantity"`
	Total_price         float32 `gorm:"column:total_price" json:"total_price"`
	VAT_percentage      uint16  `gorm:"column:vat_percentage" json:"vat_percentage"`
	Orders_Order_ID     uint64  `gorm:"column:orders_order_id" json:"orders_order_id"`
	Products_Product_ID uint32  `gorm:"column:products_product_id" json:"products_product_id"`
	Product_name        string  `gorm:"-" json:"product_name" form:"-"`
}

func NewOrderedProducts(quantity uint8, totalPrice float32, vatPercentage uint16, orderID uint64, productID uint32) *OrderedProducts {
	return &OrderedProducts{
		Quantity:            quantity,
		Total_price:         totalPrice,
		VAT_percentage:      vatPercentage,
		Orders_Order_ID:     orderID,
		Products_Product_ID: productID,
	}
}
