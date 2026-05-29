package repository

import (

	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)


func GetUser_Repo(user *[]domain.User_List ,db *gorm.DB) error{
	query := `SELECT id_user , email_user FROM public.users`
	
	err := db.Raw(query).Scan(user).Error

	if err != nil { 
		return err
	}

	return err
}


func Register_Repo(user *domain.User , db *gorm.DB ) error { 
	query := `
	INSERT INTO public.users (email_user,password_hash)
	VALUES(?,?);
	`
	err :=  db.Exec(query,user.EMAIL,user.PASSWORD).Error

   if err != nil {
	 		return err
		}

	return err
}


func GetUserByEmail_Repo(user *domain.User , db *gorm.DB) (*domain.User ,error) {
	
	user_DB := new(domain.User)

	query := `
	SELECT password_hash , id_user FROM public.users WHERE email_user = ? LIMIT 1;
	`
	err := db.Raw(query, user.EMAIL).Scan(&user_DB).Error ;
	if err != nil {
		return user_DB , err
	}
	
	return user_DB , err
}

func DeleteUserById_Repo(user_id int, db *gorm.DB) error { 
	
	query := `
	DELETE FROM public.users WHERE  id_user = ? ;
	`
	err := db.Exec(query, &user_id).Error ; 
	if err != nil { 
    return  err
	}
	
	return  err
}

func CheckUserByID_Repo(id_user int , db *gorm.DB) bool {  		
	var result bool;
	query := `
SELECT CASE 
         WHEN EXISTS(SELECT 1 FROM public.users WHERE id_user = ?) 
         THEN TRUE 
         ELSE FALSE 
       END AS is_exist;
	`
	err := db.Raw(query, &id_user).Scan(&result)

	if err != nil { 
		return result
	}

	return result
}



