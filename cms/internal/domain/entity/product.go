package entity

type Product struct {
	BaseModel
	ProductID                        uint64  `gorm:"primaryKey;autoIncrement;column:product_id" json:"ProductID" form:"-"`
	Name                             string  `gorm:"column:name" json:"name" form:"name"`
	IsAnimal                         bool    `gorm:"column:is_animal" json:"is_animal" nullable:"true"`
	Price                            float32 `gorm:"column:price" json:"price"`
	Availability                     bool    `gorm:"column:availability" json:"availability" nullable:"true"`
	VatPercentage                    float32 `gorm:"column:vat_percentage" json:"vat_percentage" form:"select" options:"/vat_percentage_dictionary" labelField:"Percentage" valueField:"Percentage"`
	ProductCode                      string  `gorm:"column:product_code" json:"product_code"`
	Description                      string  `gorm:"column:description" json:"description"`
	ProductCategory                  string  `gorm:"column:product_category" json:"product_category" form:"select" options:"/product_category_dictionary" labelField:"Name" valueField:"Name"`
	ProductMetadataProductMetadataID uint    `gorm:"-" form:"-"`
	ProductMainImage                 string  `gorm:"-" json:"product_main_image" form:"image"`
	ProductIconImage                 string  `gorm:"-" json:"product_icon_image" form:"image"`
}

func NewProduct(name, productCode, description, productCategory string, isAnimal bool, price float32, availability bool, vatPercentage float32, productMetadataID uint) *Product {
	return &Product{
		Name:                             name,
		IsAnimal:                         isAnimal,
		Price:                            price,
		Availability:                     availability,
		VatPercentage:                    vatPercentage,
		ProductCode:                      productCode,
		Description:                      description,
		ProductCategory:                  productCategory,
		ProductMetadataProductMetadataID: productMetadataID,
	}
}

func (p *Product) SetID(ID uint64) {
	p.ProductID = ID
}
