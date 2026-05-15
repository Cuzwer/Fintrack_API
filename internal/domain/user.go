package domain

type User struct { 
	ID int `gorm:"collumn:id_user" json:"id_user"` 
	EMAIL string `gorm:"collumn:email_user" json:"email"`
	PASSWORD string `gorm:"collumn:password_hash" json:"password"`
	Create_Date string `gorm:"collumn:created_at" json:"date"`
}
