package repository

import (
	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)


func PostTrans_Repo(newTrans *domain.Transaction , db *gorm.DB ) error {
	query := `
	INSERT INTO public.transactions (id_account , id_category , amount_trans , type_trans , descrip_trans )
	VALUES (?,?,?,?,?);
	`
	result := db.Exec(query, newTrans.ID_account , newTrans.ID_category , newTrans.Amount_trans , newTrans.Type_trans , newTrans.Desc_trans)

	err := result.Error

	if err != nil {
		return  err
	}

	return err
}

func GetTrans_Repo(usrId int ,newTrans *[]domain.Transaction_SentBack , db *gorm.DB ) error {
	query := `
SELECT 
	t.id_trans,
	a.name AS account,
	c.name_category AS category,
	t.amount_trans,
	t.type_trans,
	t.descrip_trans,
	t.transaction_date
FROM public.transactions t   
JOIN public.accounts a ON t.id_account = a.id_account
JOIN public.categories c ON t.id_category = c.id_category
WHERE a.id_user = ? ;
`
 err := db.Raw(query, &usrId ).Scan(newTrans).Error
	
 if err != nil {
	 return err
 }

 return nil
}
