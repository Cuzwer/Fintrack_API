package domain

import (
	"time"

	"github.com/shopspring/decimal"
)
type Transaction struct {
  ID_trans int `gorm:"column:id_trans" json:"id_trans"`
	ID_account int `gorm:"column:id_account " json:"id_account"`
	ID_category int `gorm:"column:id_catog:ory" json:"id_catogory"`
	Amount_trans decimal.Decimal `gorm:"column:amount_trans;type:decimal(15,2);not null" json:"amount_trans"`
	Type_trans string `gorm:"column:type_trans;type:varchar(20);not null" json:"type_trans"`
	Desc_trans string `gorm:"column:descrip_trans" json:"descrip_trans"`
	Transaction_Date time.Time `gorm:"column:transaction_date;type:date;not null;default:current_date" json:"transaction_date"`
  CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:current_timestamp" json:"created_at"`
}
