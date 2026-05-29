package domain


type Categories struct {
	ID_Category int `gorm:"collumn:id_category" json:"id_category"`
	ID_User int `gorm:"collumn:id_user" json:"id_user,omitempty"`
	Name_Category string `gorm:"collumn:name_category" json:"name_category"`
	Type_Category string `gorm:"collumn:type_category" json:"type_category"`
}

type Cat_Delete struct {

	ID_Category int `gorm:"collumn:id_category" json:"id_category"`
	ID_User uint `gorm:"collumn:id_user" json:"id_user,omitempty"`

}

type Cat_Action struct {
	ID_User uint `gorm:"collumn:id_user" json:"id_user,omitempty"`
	Name_Category string `gorm:"collumn:name_category" json:"name_category"`
	Type_Category string `gorm:"collumn:type_category" json:"type_category"`
}
