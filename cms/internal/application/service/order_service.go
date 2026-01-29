package service

import (
	"fmt"
	"prestaprest/internal/domain/entity"
	"prestaprest/internal/infrastructure/postgres"

	"gorm.io/gorm"
)

type OrderService struct {
	PostgresRepository        postgres.PostgresRepository[entity.Order]
	OrderedProductsRepository postgres.PostgresRepository[entity.OrderedProducts]
	ProductRepository         postgres.PostgresRepository[entity.Product]
}

func NewOrderService(conn *gorm.DB) OrderService {
	return OrderService{
		PostgresRepository:        postgres.NewPostgresRepository[entity.Order](conn),
		OrderedProductsRepository: postgres.NewPostgresRepository[entity.OrderedProducts](conn),
		ProductRepository:         postgres.NewPostgresRepository[entity.Product](conn),
	}
}

func (s *OrderService) GetAllOrders() ([]entity.Order, error) {
	allOrders, err := s.PostgresRepository.GetAll("Order_ID")
	if err != nil {
		return nil, err
	}
	return allOrders, nil
}

func (s *OrderService) GetAllClientOrders(client_ID uint64) ([]entity.Order, error) {
	allOrders, err := s.PostgresRepository.GetAllFiltered("Order_ID", "clients_client_id", client_ID)
	if err != nil {
		return nil, err
	}
	return allOrders, nil
}

func (s *OrderService) GetAllOrderProducts(clientID uint64) ([]entity.Order, error) {
	loadProducts := func(order *entity.Order) error {
		type OrderedProductWithName struct {
			Quantity            uint8
			Total_price         float32
			VAT_percentage      uint16
			Orders_Order_ID     uint64
			Products_Product_ID uint32
			Product_name        string
		}

		var tempResults []OrderedProductWithName

		err := s.OrderedProductsRepository.GetDB().
			Table("ordered_products").
			Select("ordered_products.quantity, ordered_products.total_price, ordered_products.vat_percentage, ordered_products.orders_order_id, ordered_products.products_product_id, products.name as product_name").
			Joins("LEFT JOIN products ON products.product_id = ordered_products.products_product_id").
			Where("ordered_products.orders_order_id = ?", order.Order_ID).
			Scan(&tempResults).Error

		if err != nil {
			return err
		}

		orderedProducts := make([]entity.OrderedProducts, len(tempResults))
		for i, temp := range tempResults {
			orderedProducts[i] = entity.OrderedProducts{
				Quantity:            temp.Quantity,
				Total_price:         temp.Total_price,
				VAT_percentage:      temp.VAT_percentage,
				Orders_Order_ID:     temp.Orders_Order_ID,
				Products_Product_ID: temp.Products_Product_ID,
				Product_name:        temp.Product_name,
			}
		}

		order.OrderedProducts = orderedProducts
		return nil
	}

	return s.PostgresRepository.GetAllWithRelation("clients_client_id", clientID, "order_id", loadProducts)
}

func (s *OrderService) GetOrder(orderID uint64) (*entity.Order, error) {
	order, err := s.PostgresRepository.Get("Order_ID", orderID)
	fmt.Println("Am here (service)?")
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) CreateOrder(order *entity.Order) error {
	err := s.PostgresRepository.Create(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) CreateOrderProduct(cart map[uint64]int, order_ID uint64) (float32, error) {
	fmt.Println(cart)
	var totalPrice float32 = 0
	for item, quantity := range cart {
		product, err := s.ProductRepository.Get("Product_ID", item)
		if err != nil {
			fmt.Println("Error getting product:", item, err)
			return 0, err
		}
		orderItem := entity.NewOrderedProducts(
			uint8(quantity),
			product.Price*float32(quantity),
			uint16(product.VatPercentage),
			order_ID,
			uint32(product.ProductID),
		)
		totalPrice += product.Price * float32(quantity)
		err = s.OrderedProductsRepository.Create(orderItem)
		if err != nil {
			return 0, err
		}
	}
	return totalPrice, nil
}

func (s *OrderService) UpdateOrder(order *entity.Order) error {
	err := s.PostgresRepository.Update(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) DeleteOrder(order_ID uint) error {
	err := s.PostgresRepository.Delete("Order_ID", uint64(order_ID))
	if err != nil {
		return err
	}
	return nil
}
