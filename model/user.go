package model

type User struct {
	BaseModel
	CreatedBy string `gorm:"default:admin" json:"-"`
	UpdatedBy string `gorm:"default:admin" json:"-"`
	Email     string `gorm:"unique" `
	Password  string `json:"-"`
	IsActive  bool   `gorm:"default:false"`
	FirstName string
	LastName  string
	Roles     []Role `gorm:"many2many:users_roles"`
}
