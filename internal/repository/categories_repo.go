package repository

import (

	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)


type Categories struct {
	ID_Category int `gorm:"collumn:id_category" json:"id_category"`
	ID_User int `gorm:"collumn:id_user" json:"id_user,omitempty"`
	Name_Category string `gorm:"collumn:name_category" json:"name_category"`
	Type_Category string `gorm:"collumn:type_category" json:"type_category"`
}


func GetCategory(id_user uint ,catagory *[]domain.Categories , db *gorm.DB) error { 
	query := `
	SELECT id_category , name_category , type_category 
	FROM public.categories
	WHERE id_user = ?;
	`
	err := db.Raw(query,&id_user).Scan(catagory).Error
	
	if err != nil {
		return err
	}

	return err
}

func DeleteCategory_Repo(cat *domain.Cat_Delete , db *gorm.DB)  error {

	query := `
	DELETE FROM public.categories WHERE id_category = ? AND id_user = ?;
	`
	result := db.Exec(query, cat.ID_Category , cat.ID_User)
	err := result.Error

	if err != nil  { 
		return  err
	}
	return  err
}


func PostCat_Repo(cat *domain.Cat_Action , db *gorm.DB ) error {
	query := `
	INSERT INTO public.categories (id_user ,name_category ,  type_category )
	VALUES (? , ? , ? );
	`
	
	err := db.Exec(query, cat.ID_User ,cat.Name_Category , cat.Type_Category).Error

	if err != nil {
		return  err

	}
	
	return err;
}

func UpdateCat_repo(cat *domain.Categories , db *gorm.DB ) error { 
	query := `
	UPDATE public.categories
	SET  name_category = ?,
	type_category = ?
	WHERE id_user = ? AND id_category = ?;
	`
	err := db.Exec(query, cat.Name_Category , cat.Type_Category , cat.ID_User , cat.ID_Category).Error

 	if err != nil {

		return  err
	}
	return  err
} 

func CheckCatByuserId_Repo(usrId int , CatId int , db *gorm.DB ) error { 
	var result bool 

	query := `
	SELECT EXISTS(SELECT 1 FROM public.categories WHERE id_user = ? AND id_category = ?)
	`
	err := db.Raw(query ,usrId ,CatId).Scan(&result).Error

	if err != nil || result != true { 
		return err
	}

	return nil;
}
