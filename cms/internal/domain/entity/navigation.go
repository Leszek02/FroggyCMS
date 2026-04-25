package entity

type Navigation struct {
	BaseModel
	Navigation_ID                    uint          `gorm:"primaryKey;column:navigation_id" json:"navigation_id" form:"-"`
	Name                             string        `gorm:"column:name" json:"name"`
	Url                              string        `gorm:"column:url" json:"url"`
	Position                         int           `gorm:"column:position" json:"position"`
	Status                           bool          `gorm:"column:status" json:"status" nullable:"true"`
	Navigation_Navigation_ID         *uint         `gorm:"column:navigation_navigation_id" json:"parent" form:"select" options:"/navigation" labelField:"Name" valueField:"Navigation_ID"`
	Administrators_Administrators_ID uint          `gorm:"column:administrators_administrator_id" json:"administrator_id" form:"-"`
	Children                         []*Navigation `gorm:"-" json:"children,omitempty" form:"-"`
}

func NewNavigation(name, url string, status bool, position int, navigation_parent_ID *uint, administrator_ID uint) *Navigation {
	return &Navigation{
		Name:                             name,
		Url:                              url,
		Position:                         position,
		Status:                           status,
		Navigation_Navigation_ID:         navigation_parent_ID,
		Administrators_Administrators_ID: administrator_ID,
	}
}

func (n *Navigation) SetID(ID uint) {
	n.Navigation_ID = ID
}

func (Navigation) TableName() string {
	return "navigation"
}
