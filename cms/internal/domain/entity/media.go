package entity

type Media struct {
	BaseModel
	Media_ID uint64 `gorm:"primaryKey;autoIncrement"`
	Name     string `form:"Name" json:"name"`
	Path     string `form:"Path" json:"path"`
}

func NewMedia(name, path string) *Media {
	return &Media{
		Name: name,
		Path: path,
	}
}

func (m *Media) SetID(ID uint64) {
	m.Media_ID = ID
}

func (m *Media) TableName() string {
	return "medias"
}

type Products_Medias struct {
	Products_Product_ID uint64
	Medias_Media_ID     uint64
}

func NewProductsMedias(productID, mediaID uint64) *Products_Medias {
	return &Products_Medias{
		Products_Product_ID: productID,
		Medias_Media_ID:     mediaID,
	}
}
