package handler

import (
	"strconv"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryOBG_handeler struct { 
  DB *gorm.DB
}

func CategoryHandler(db *gorm.DB) *CategoryOBG_handeler{
	
	var cat_obg CategoryOBG_handeler
	cat := &cat_obg
	cat.DB = db

	return cat
}


func(h *CategoryOBG_handeler) GetCategory() fiber.Handler{
	return  func(c *fiber.Ctx)  error {
		val  := c.Locals("userID")
		userID , ok := val.(uint)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "User ID format is invalid ",
			})	
		}
		var CategoriesOBJ []domain.Categories
		cat := &CategoriesOBJ

		err := service.GetCategories_service(userID , cat , h.DB )

		if err != nil { 
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Something went wrong cannot Get Category ",
			})	
		}

		return  c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"status" : fiber.StatusAccepted,
			"result" : cat,
		})	
	}
}

func(h* CategoryOBG_handeler) DeleteCat_handler() fiber.Handler {
	return func (c *fiber.Ctx) error {
		val := c.Locals("userID")
		userID , ok := val.(uint)
		if !ok { 
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "User ID format invalid",
				"userID" : userID,
			})
		}
    id_cat , err := strconv.Atoi(c.Params("id"))

		if err != nil { 
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Id catagory invalid",
			})
		}

		var cat_detail domain.Cat_Delete;
		cat := &cat_detail
		cat.ID_User = userID
		cat.ID_Category = id_cat
		
		err = service.DeleteCat_service(cat , h.DB )	
		
		if err != nil {
 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
    		"message" : "failed to deleted category please try another id category",
			})
		}
		
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "OK",
      "value" : cat,
		})
	}
}

func(h * CategoryOBG_handeler) PostCat_hanler() fiber.Handler { 
	return  func(c *fiber.Ctx) error {
		val := c.Locals("userID")
		userID , ok := val.(uint)

		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "userID Format invalid",
				"userID" : userID,
			})
		}
		var catPost domain.Cat_Action
		cat := &catPost
		err := c.BodyParser(cat)
		
		if err != nil { 
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "invalid format json",
			})
		}
		cat.ID_User = userID;

		err = service.PostCat_service(cat ,h.DB)	
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message" : "faild to create Categories Something went wrong",
				"error" : err,
			})	
		}
		
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "Sucessfully for Post Categories",
		})
	}
}


func (h *CategoryOBG_handeler) UpdateCat_handler() fiber.Handler {
	return  func(c *fiber.Ctx) error { 
	  val := c.Locals("userID")
		userID , ok := val.(uint)
		
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("invalid format user id")
		}
	  
		idCat , err := strconv.Atoi(c.Params("id"))
   
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Invalid format Categories id ")
		}

		var cat_obj domain.Categories
		cat := &cat_obj
		
		if err := c.BodyParser(cat) ; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Invalid Input format",
			})
		}

		cat.ID_User = int(userID)
    cat.ID_Category = idCat

		if err := service.UpdateCat_service(cat , h.DB) ; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "faild to updated category",
			})
		}
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "Sucessfully updated category",
			"status code" : fiber.StatusAccepted,
		})
	}
}
