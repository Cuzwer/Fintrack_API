package service

import (
	"fmt"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"
)


func GetCategories_service(userID uint ,cat *[]domain.Categories , db *gorm.DB) error { 
	
	if err := repository.CheckUserByID_Repo(int(userID), db ) ; err != true {
		return fmt.Errorf("\n there 's  no this userID in the database %v \n" , err )
	}

	err := repository.GetCategory(userID , cat , db )

	if err != nil { 
		return fmt.Errorf("\n Something went wrong with get Categories %v \n", err)
	}

	return  err 
}

func DeleteCat_service(cat *domain.Cat_Delete , db *gorm.DB) error { 
	
	if err := repository.CheckUserByID_Repo(int(cat.ID_User) , db ) ; err != true {
		return fmt.Errorf("\n there 's  no this userID in the database %v \n" , err )
	}

	err  := repository.DeleteCategory_Repo(cat ,db)

	if err != nil {
		return fmt.Errorf("\n Something went wrong fialed to Delete Categories %v", err)
	}

	return  err
}

func PostCat_service(cat *domain.Cat_Action , db *gorm.DB) error {


	if cat.Type_Category == "" || cat.Name_Category == "" {
		return fmt.Errorf("\n request body is nil Failed to post \n")
	}

	err := repository.PostCat_Repo(cat , db)
	
	if err != nil { 
		return fmt.Errorf("\n Failed Post Categories %v", err)
	}

	return err
}

func UpdateCat_service(cat *domain.Categories , db *gorm.DB) error { 
	
	if cat.Type_Category == "" || cat.Name_Category == ""  {
		return fmt.Errorf("\n request body is nil Failed to post \n")
	}
	
  err := repository.UpdateCat_repo(cat , db)

	if err != nil {
		return fmt.Errorf("\n Failed to Update Category %v", err)
	}
	return  err;

}


func CheckCatByuserId_Service(userID int , idCat int  , db *gorm.DB ) error { 
	if userID == 0 || idCat == 0  { 
		return  fmt.Errorf("userID and id cat is zero");
	}

	err := repository.CheckCatByuserId_Repo(userID , idCat, db)

	if err != nil { 
		return  fmt.Errorf("userID and id category did not link to each other");
	}

	return nil;
}
