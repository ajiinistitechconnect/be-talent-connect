package model

type User struct {
	BaseModel
	CreatedBy string
	UpdatedBy string
	Email     string
	Password  string
	IsActive  bool
	FirstName string
	LastName  string
	Roles     []Role `gorm:"many2many:users_roles"`
}
