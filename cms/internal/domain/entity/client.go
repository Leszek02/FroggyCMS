package entity

import "time"

type Client struct {
	BaseModel
	Client_ID  uint64    `gorm:"primaryKey;autoIncrement" form:"-"`
	First_name string    `form:"First_name" json:"First_name"`
	Last_name  string    `form:"Last_name" json:"Last_name"`
	Login      string    `form:"login" json:"login"`
	Password   string    `json:"password" form:"-"`
	Email      string    `form:"email" json:"email"`
	Phone      string    `form:"phone" json:"phone"`
	Status     bool      `form:"status" json:"status" nullable:"true"`
	Newsletter bool      `form:"newsletter" json:"newsletter" nullable:"true"`
	Last_login time.Time `form:"last_login" json:"last_login"`
	CreatedAt  time.Time `gorm:"column:join_date;not null"`
}

func NewClient(firstName, lastName, login, password, email, phone string) *Client {
	return &Client{
		First_name: firstName,
		Last_name:  lastName,
		Login:      login,
		Password:   password,
		Email:      email,
		Phone:      phone,
		Status:     true,
		Newsletter: false,
	}
}

func (c *Client) SetID(ID uint64) {
	c.Client_ID = ID
}
