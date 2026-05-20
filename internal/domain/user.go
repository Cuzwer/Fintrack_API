package domain

type User struct { 
	ID int `gorm:"column:id_user;primaryKey;autoIncrement" json:"id_user"` 
	EMAIL string `gorm:"column:email_user;uniqe;not null" json:"email"`
	PASSWORD string `gorm:"column:password_hash;not null" json:"password"`
	Create_Date string `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

type User_List struct { 
	ID int `gorm:"column:id_user;primaryKey;autoIncrement" json:"id_user"` 
	EMAIL string `gorm:"column:email_user;uniqe;not null" json:"email"`
}


type User_Login struct {
	EMAIL string `gorm:"column:email_user;uniqe;not null" json:"email"`
	PASSWORD string `gorm:"column:password_hash;not null" json:"password"`
	JWT string `json:"jwt"`
}
