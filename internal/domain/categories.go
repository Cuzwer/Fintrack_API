package domain


type Categories struct {
	ID_Category int `gorm:"collumn:id_category" json:"id_category"`
	ID_User int `gorm:"collumn:id_user" json:"id_user"`
	Name_Category string `gorm:"collumn:name_category" json:"name_category"`
	Type_Category string `gorm:"collumn:type_category" json:"type_category"`
}
