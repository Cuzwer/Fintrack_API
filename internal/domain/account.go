package domain

import "github.com/shopspring/decimal"


type Account struct {
	ID_Account int `gorm:"collum:id_account;primaryKey;autoIncrement" json:"id_account"`
	ID_User int `gorm:"collumn:id_user" json:"id_user"`
	Name string `gorm:"collumn:name" json:"name"`
	TYPE string `gorm:"column:type_account" json:"type"`
 	Balance decimal.Decimal `gorm:"collumn:balance" json:"balance"`
	Currency string `gorm:"collumn:currency" json:"currency"`
}


type Detail_Account struct {
	ID_User int `gorm:"collumn:id_user" json:"id_user"`
	Name string `gorm:"collumn:name" json:"name"`
	TYPE string `gorm:"column:type_account" json:"type"`
 	Balance decimal.Decimal `gorm:"collumn:balance" json:"balance"`
	Currency string `gorm:"collumn:currency" json:"currency"`
}
