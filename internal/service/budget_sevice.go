package service

import (
	"fmt"
	"time"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"
)



func PostBudget_service(newBudget *domain.Budgets, db *gorm.DB ) error { 
		err := CheckCatByuserId_Service(newBudget.ID_user ,  newBudget.ID_category , db )

		if err != nil { 
			return  fmt.Errorf("userID and id category did not link to each other");
		}
		
		if err := newBudget.Amount_budg.IsZero() ; err != false {
			return  fmt.Errorf("\n Amount Budgets can't be zero \n")
		}

		now := time.Now()
		if newBudget.Month_budg == 0  || newBudget.Year_budg == 0 { 
			newBudget.Month_budg = int(now.Month())
		  newBudget.Year_budg = int(now.Year());
		}

		err = repository.PostBudg_Repo(newBudget  ,  db);
		if err != nil { 
			return  fmt.Errorf("failed to post budget");
		}

		return nil
}

func GetBudget_Service(userID int, allbudg *[]domain.Budgets_detail , db *gorm.DB) error { 
		
	err := repository.GetBudget_Repo(userID ,  allbudg , db ) 

	if err  != nil {
		return  fmt.Errorf("\n failed to get budger for this user \n")
	}

	return  nil;
}

func UpdateBudg_Service(budg *domain.Budgets_Edit ,db *gorm.DB) error { 
	
	if budg.ID_category == 0 || budg.Amount_budg.IsZero() != false  {
		return fmt.Errorf("\n Invalid Input i dont know which but if could be category and amount \n");
	}

	if  err := CheckCatByuserId_Service( budg.ID_user ,  budg.ID_category , db) ; err != nil {
		return  fmt.Errorf("userID and id category did not related to each other");
	}

	now := time.Now()

	if budg.Year_budg == 0 || budg.Month_budg == 0 {
		budg.Year_budg = now.Year()
		budg.Month_budg = int(now.Month())
	}
	
	err := repository.UpdateBudget_Repo(budg , db)	

	if err != nil {
		return fmt.Errorf("Something went wrong can't update budget");
	}
	return nil
}

func DeleteBudge_Service(detail *domain.Budget_Delete , db *gorm.DB) error {
	if detail.Id_bug == 0 || detail.ID_user == 0 {
		return fmt.Errorf("userID or id budg Invalid");
	}
	
	err := repository.DeleteBudget_Repo(detail , db);

	if err != nil {
		return fmt.Errorf("failed to deleted something went wrong");
	}

	return nil
}
