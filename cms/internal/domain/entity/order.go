package entity

type Order struct {
	BaseModel
	Order_ID          uint64  `gorm:"primaryKey;autoIncrement"`
	Status            bool    `form:"Status" json:"Status"`
	Payment_name      string  `form:"Payment_name" json:"Payment_name"`
	Payment_price     float32 `form:"Payment_price" json:"Payment_price"`
	Delivery_name     string  `form:"Delivery_name" json:"Delivery_name"`
	Delivery_price    float32 `form:"Delivery_price" json:"Delivery_price"`
	Order_note        string  `form:"Order_note" json:"Order_note"`
	Address_city      string  `form:"address_city" json:"address_city"`
	Address_postcode  string  `form:"Address_postcode" json:"Address_postcode"`
	Address_street    string  `form:"Address_street" json:"Address_street"`
	Address_state     string  `form:"Address_state" json:"Address_state"`
	Total_price       float32 `form:"Total_price" json:"Total_price"`
	Clients_Client_ID uint64
	OrderedProducts   []OrderedProducts `gorm:"-" form:"-"`
}

func NewOrder(status bool,
	paymentName string,
	paymentPrice float32,
	deliveryName string,
	deliveryPrice float32,
	orderNote string,
	addressCity string,
	addressPostcode string,
	addressStreet string,
	addressState string,
	totalPrice float32,
	clientsClientID uint64) *Order {
	return &Order{
		Status:            status,
		Payment_name:      paymentName,
		Payment_price:     paymentPrice,
		Delivery_name:     deliveryName,
		Delivery_price:    deliveryPrice,
		Order_note:        orderNote,
		Address_city:      addressCity,
		Address_postcode:  addressPostcode,
		Address_street:    addressStreet,
		Address_state:     addressState,
		Total_price:       totalPrice,
		Clients_Client_ID: clientsClientID,
	}
}

func (o *Order) SetID(id uint) {
	o.Order_ID = uint64(id)
}

func (o *Order) SetClientID(id uint64) {
	o.Clients_Client_ID = id
}

func (o *Order) UpdateStatus(status bool) {
	o.Status = status
}

func (o *Order) SetPaymentMethod(payment *PaymentDictionary) {
	o.Payment_name = payment.Name
	o.Payment_price = payment.Price
}

func (o *Order) SetDeliveryMethod(delivery *DeliveryDictionary) {
	o.Delivery_name = delivery.Name
	o.Delivery_price = delivery.Price
}
