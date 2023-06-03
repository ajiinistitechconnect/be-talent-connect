package model

type User struct {
	BaseModel
	CreatedBy string `gorm:"default:admin"`
	UpdatedBy string `gorm:"default:admin"`
	Email     string `gorm:"unique"`
	Password  string `json:"-"`
	IsActive  bool   `gorm:"default:false"`
	FirstName string
	LastName  string
	Roles     []Role `gorm:"many2many:users_roles"`
}
