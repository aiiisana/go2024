package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" gorm:"size:255"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Tasks    []Task `gorm:"foreignKey:UserID"`
}
