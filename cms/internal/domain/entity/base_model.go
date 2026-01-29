package entity

import "time"

type BaseModel struct {
	Create_date time.Time `gorm:"column:create_date;autoCreateTime" json:"create_date" form:"text"`
	Modify_date time.Time `gorm:"column:modify_date;autoUpdateTime" json:"modify_date" form:"text"`
}
