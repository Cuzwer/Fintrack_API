package domain


type Account struct {
	ID_Account int `gorm:"collum:id_account;primaryKey;autoIncrement" json:"id_account"`
	ID_User int `gorm:"collumn:id_user" json:"id_user"`
	Name string `gorm:"collumn:name" json:"name"`
 	Type string `gorm:"collumn:balance" json:"balance"`
	Currency string `gorm:"collumn:currency" json:"currency"`
}

