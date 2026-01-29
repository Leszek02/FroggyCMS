package entity

type Page struct {
	BaseModel
	Page_ID                         uint   `gorm:"primaryKey;autoIncrement;column:page_id" json:"page_id" form:"page_id"`
	Name                            string `gorm:"column:name" json:"name" form:"name"`
	Type                            string `gorm:"column:type" json:"type" form:"type"`
	Content                         string `gorm:"column:content" json:"content" form:"content"`
	Status                          bool   `gorm:"column:status" json:"status" nullable:"true" form:"status"`
	Administrators_Administrator_ID int    `gorm:"column:administrators_administratod_id" json:"administrator_id"`
}

func NewPage(name, page_type, content string, status bool, administrator int) *Page {
	return &Page{
		Name:                            name,
		Type:                            page_type,
		Content:                         content,
		Status:                          status,
		Administrators_Administrator_ID: administrator,
	}
}
