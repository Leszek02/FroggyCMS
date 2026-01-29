package entity

type Administrator struct {
	BaseModel
	Administrator_ID uint32 `gorm:"primaryKey;autoIncrement" form:"-"`
	First_name       string `form:"First_name" json:"First_name"`
	Last_name        string `form:"Last_name" json:"Last_name"`
	Login            string `form:"Login" json:"Login"`
	Password         string `form:"Password" json:"Password"`
	Email            string `form:"Email" json:"Email"`
	Phone            string `form:"Phone" json:"Phone"`
}

func NewAdministrator(firstName, lastName, login, password, email, phone string) *Administrator {
	return &Administrator{
		First_name: firstName,
		Last_name:  lastName,
		Login:      login,
		Password:   password,
		Email:      email,
		Phone:      phone,
	}
}

func (a *Administrator) SetID(id uint) {
	a.Administrator_ID = uint32(id)
}
