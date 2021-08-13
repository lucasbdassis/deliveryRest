package models

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required" json:"user" gorm:"column:user"`
	Password string `form:"password" json:"password" xml:"password" binding:"required" json:"password" gorm:"column:password"`
}
