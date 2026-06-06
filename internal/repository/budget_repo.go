package repository

import (

	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)


func PostBudg_Repo(newBudg *domain.Budgets , db *gorm.DB)  error {
	query := `
	INSERT INTO public.budgets (id_user , id_category , amount_budg , month_budg , year_budg)
	VALUES (? , ?, ?, ?, ?);
	`
	
	err := db.Exec(query , newBudg.ID_user , newBudg.ID_category , newBudg.Amount_budg , newBudg.Month_budg ,  newBudg.Year_budg).Error ; 
	
	if err != nil { 
		return err;
	}

	return  nil
}

func GetBudget_Repo(usrID int, AllBudg *[]domain.Budgets_detail, db *gorm.DB ) error { 
	query := `
	SELECT
	b.id_budg,
	b.id_user,
	c.name_category as category,
	b.amount_budg,
	b.month_budg,
	b.year_budg
	FROM public.budgets b
	JOIN public.categories c ON b.id_category = c.id_category
	WHERE b.id_user = ?;
	`

	if err := db.Raw(query , usrID ).Scan(AllBudg).Error ; err != nil {
		return  err
	}

	return nil;
}

func UpdateBudget_Repo(budg *domain.Budgets_Edit ,  db *gorm.DB) error { 
	query := `
	UPDATE public.budgets
	SET id_category = ?,
	amount_budg = ? ,
	month_budg = ? ,
	year_budg = ? 
	WHERE id_budg = ?;
	`

	err := db.Exec(query, budg.ID_category , budg.Amount_budg , budg.Month_budg ,  budg.Year_budg , budg.Id_bug).Error

	if err != nil { 
		return err
	}

	return nil
}

func DeleteBudget_Repo(budg *domain.Budget_Delete , db *gorm.DB) error {
	query  := `
	DELETE FROM public.budgets
	WHERE id_budg = ? AND id_user = ? ;
	`

	err := db.Exec(query , budg.Id_bug , budg.ID_user).Error
	if err != nil {
		return err
	}
	return nil
}
