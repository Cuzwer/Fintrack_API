package service

import (
	"fmt"
	"log"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"
)



func PostTransaction(ID_User uint, newTran *domain.Transaction , db *gorm.DB ) error {
	if newTran.ID_account == 0  || newTran.ID_category == 0 || newTran.Amount_trans.IsZero() {
		return fmt.Errorf("\n invalid input: missing account, category, or amount \n")
	}

	if newTran.Type_trans == "" {
		return fmt.Errorf("\n invalid type transaction \n")
	}
	account_check := new(domain.AccountCheckUser)
	account_check.ID_Account = newTran.ID_account
	account_check.ID_User = int(ID_User)

	if err := CheckAccountByuserID_Service(account_check, db) ; err != nil {
		log.Printf("Account verification failed: %v", err)

		return  fmt.Errorf("account verification failed: %w", err)
	}
	account_buffer := new(domain.Account);
	account_buffer.ID_Account = newTran.ID_account

	balance := repository.CheckMoney_ByIDAccount_Repo(&account_check.ID_Account ,( db))

	if newTran.Type_trans == "expense" {

		if newTran.Amount_trans.GreaterThan(balance) {
				return fmt.Errorf("Transaction failed: balance is less than transaction amount")
		} else { 
			account_buffer.Balance = balance.Sub(newTran.Amount_trans);
 			ChangeMoneyByAcc_Service(account_buffer , db)
		}
	} else {
		account_buffer.Balance = balance.Add(newTran.Amount_trans);
		
			ChangeMoneyByAcc_Service(account_buffer, db);
	}

	err := repository.PostTrans_Repo(newTran , db )
	if err != nil {
		return fmt.Errorf("\n Cant Post Transaction Some thing went wrong \n ")
	}

	return nil
	}

func GetTransaction_Service(usrID int , Trans *[]domain.Transaction_SentBack , db *gorm.DB) error {


	 err := repository.GetTrans_Repo(usrID, Trans , db)

	 if err != nil {
		 return fmt.Errorf("failed to get transaction \n")
	 }

	 return nil
}
