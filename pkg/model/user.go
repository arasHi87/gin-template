package model

type UserModel struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name;not null"`
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
}
