package repository

import (
	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)



func GetAll_AccountUser_Repo(id_user  int ,account_all *[]domain.Detail_Account, db *gorm.DB)	error { 
	query := `
	SELECT id_account ,id_user , name , type_account , balance , currency FROM public.accounts
	WHERE id_user = ?;
	`

	err := db.Raw(query,&id_user).Scan(account_all).Error

	if err != nil { 
		return err
	}

	return err;
} 


func PostAccount_Repo (NewAccount *domain.Detail_Account , db *gorm.DB ) error { 
	query := `
	INSERT INTO public.accounts (id_user ,name , type_account , balance , currency )
	VALUES (?,?,?,?,?)
	`
	
	result := db.Exec(query, NewAccount.ID_User ,NewAccount.Name , NewAccount.TYPE , NewAccount.Balance , NewAccount.Currency)
	rowAff := result.RowsAffected
	err := result.Error

	if err != nil || rowAff != 1 {
		return err
	}
	return  err
}

func CheckAccountByIduser_Repo (account *domain.AccountCheckUser , db *gorm.DB) (error,  bool){
	var exists bool;

	query := `
	SELECT EXISTS(SELECT 1 FROM public.accounts WHERE id_account = ? AND  id_user = ?)
	`
	err := db.Raw(query, account.ID_Account , account.ID_User ).Scan(&exists).Error
	
	if err != nil { 
		return err , exists
	}
	return err , exists
}
