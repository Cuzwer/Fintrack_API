package domain

import "github.com/shopspring/decimal"

type Budgets struct {
	Id_bug int `gorm:"column:id_budg;primaryKey;autoIncrement" json:"id_budg"`
	ID_user int `gorm:"column:id_user" json:"id_user"`
 	ID_category int `gorm:"column:id_category" json:"id_category"`
	Amount_budg decimal.Decimal `gorm:"column:amount_budg;type:decimal(15,2);not null" json:"amount_budg"`
	Month_budg  int `gorm:"column:month_budg;not null" json:"month_budg"`
	Year_budg int `gorm:"column:year_budg;not null" json:"year_budg"`
}

type Budgets_detail struct {
	Id_bug int `gorm:"column:id_budg;primaryKey;autoIncrement" json:"id_budg"`
	ID_user int `gorm:"column:id_user" json:"id_user"`
 	Category string `gorm:"column:category" json:"category"`
	Amount_budg decimal.Decimal `gorm:"column:amount_budg;type:decimal(15,2);not null" json:"amount_budg"`
	Month_budg  int `gorm:"column:month_budg;not null" json:"month_budg"`
	Year_budg int `gorm:"column:year_budg;not null" json:"year_budg"`
}

type Budgets_Edit struct {
	Id_bug int `gorm:"column:id_budg;primaryKey;autoIncrement" json:"id_budg"`
	ID_user int `gorm:"column:id_user" json:"id_user"`
 	ID_category int `gorm:"column:id_category" json:"id_category"`
	Amount_budg decimal.Decimal `gorm:"column:amount_budg;type:decimal(15,2);not null" json:"amount_budg"`
	Month_budg  int `gorm:"column:month_budg;not null" json:"month_budg"`
	Year_budg int `gorm:"column:year_budg;not null" json:"year_budg"`
}



type Budget_Delete struct {
	Id_bug int `gorm:"column:id_budg;primaryKey;autoIncrement" json:"id_budg"`
	ID_user int `gorm:"column:id_user" json:"id_user"`
}
