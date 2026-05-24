package service

import (
	"fmt"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

func GetAlluser_service(user *[]domain.User_List, db *gorm.DB) error {
	return repository.GetUser_Repo(user, db)
}

func RegisterUser_service(user *domain.User, db *gorm.DB) error {
	if user.PASSWORD != " " {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.PASSWORD), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.PASSWORD = string(hashPassword)
	}

	err := repository.Register_Repo(user, db)
	if err != nil {
		return err
	}

	return err
}

func LoginUser_Service(user *domain.User , db *gorm.DB) error { 

	if user.PASSWORD == " " && user.EMAIL ==  " " {
		return fmt.Errorf("email or password cannot be empty")
	}

	  user_DB ,err := repository.GetUserByEmail_Repo(user,db)

		if err !=  nil { 
		fmt.Printf("There's no email in database please Registeration [%v] \n " , err)
		return err  }

		err = bcrypt.CompareHashAndPassword([]byte(user_DB.PASSWORD), []byte(user.PASSWORD)); 
		if err != nil { 
				return  err
		}
		user.ID = user_DB.ID
	fmt.Printf("There's email in the database and the PASSWORD is correctly [%v]", err )

	return err
}

func DeleteUserById_service(user_id int, db *gorm.DB) error { 
	
	if repository.CheckUserByID_Repo(user_id,db) != true {
		return 	fmt.Errorf("There's no this user id in the database ❌ \n");
	}

	err := repository.DeleteUserById_Repo(user_id , db)
	if err != nil { 
		return  err
	}
	
	return err;
}


