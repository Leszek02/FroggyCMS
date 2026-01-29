package entity

type DeliveryDictionary struct {
	Delivery_dict_ID uint16  `gorm:"primaryKey;autoIncrement" json:"Delivery_dict_ID" form:"-"`
	Name             string  `form:"Name" json:"Name"`
	Price            float32 `form:"Price" json:"Price"`
}

func NewDeliveryDictionary(name string, price float32) *DeliveryDictionary {
	return &DeliveryDictionary{
		Name:  name,
		Price: price,
	}
}

func (d *DeliveryDictionary) SetID(ID uint16) {
	d.Delivery_dict_ID = ID
}

func (DeliveryDictionary) TableName() string {
	return "delivery_dictionary"
}

/////////////////////////////////////////

type PaymentDictionary struct {
	Payment_dict_ID uint16  `gorm:"primaryKey;autoIncrement" form:"-"`
	Name            string  `form:"Name" json:"Name"`
	Price           float32 `form:"Price" json:"Price"`
}

func NewPaymentDictionary(name string, price float32) *PaymentDictionary {
	return &PaymentDictionary{
		Name:  name,
		Price: price,
	}
}

func (d *PaymentDictionary) SetID(ID uint16) {
	d.Payment_dict_ID = ID
}

func (PaymentDictionary) TableName() string {
	return "payment_dictionary"
}

/////////////////////////////////////////

type VATPercentageDict struct {
	VAT_percentage_dict_ID uint8  `gorm:"primaryKey;autoIncrement" form:"-"`
	Percentage             uint16 `form:"Percentage" json:"Percentage"`
}

func NewVATPercentageDict(percentage uint16) *VATPercentageDict {
	return &VATPercentageDict{
		Percentage: percentage,
	}
}

func (v *VATPercentageDict) SetID(ID uint8) {
	v.VAT_percentage_dict_ID = ID
}

func (VATPercentageDict) TableName() string {
	return "vat_percentage_dictionary"
}

/////////////////////////////////////////

type ProductCategoryDictionary struct {
	Product_dict_ID uint8  `gorm:"primaryKey;autoIncrement" form:"-"`
	Name            string `form:"Name" json:"Name"`
}

func NewProductCategoryDictionary(name string) *ProductCategoryDictionary {
	return &ProductCategoryDictionary{
		Name: name,
	}
}

func (p *ProductCategoryDictionary) SetID(ID uint8) {
	p.Product_dict_ID = ID
}

func (ProductCategoryDictionary) TableName() string {
	return "product_category_dictionary"
}

/////////////////////////////////////////
