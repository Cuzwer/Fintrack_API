package service

import (
	"fmt"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"
)

func GetAccountAll_Service( id_user int,all_account *[]domain.Detail_Account,db *gorm.DB) error {

	if err := repository.CheckUserByID_Repo(id_user, db)	; err !=  true {
		return fmt.Errorf("There's no user id in Database ")
	}	

	err := repository.GetAll_AccountUser_Repo(id_user,all_account,db);

	if err != nil {
		return err
	}

	return err
}

func PostAccount_Service(Newaccount *domain.Detail_Account, db *gorm.DB) error { 

	err := repository.PostAccount_Repo(Newaccount,db) ; 
	if err != nil { 
			return fmt.Errorf("Some thing went wrong cant Post")
	}
	
	return err
}

func CheckAccountByuserID_Service(account *domain.AccountCheckUser ,  db *gorm.DB ) error {
	err , exist :=	repository.CheckAccountByIduser_Repo(account , db)
	if err != nil || exist != true { 
		return fmt.Errorf("\n There's no account or user In account table \n")
	}
	return err
}
