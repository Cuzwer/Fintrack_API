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
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message" : "Invalid user ID",
			})	
		}
		var CategoriesOBJ []domain.Categories
		cat := &CategoriesOBJ

		err := service.GetCategories_service(userID , cat , h.DB )

		if err != nil { 
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Failed to retrieve categories",
				"error": err.Error(),
			})	
		}

		return  c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message" : "Categories retrieved successfully",
			"result" : cat,
		})	
	}
}

func(h* CategoryOBG_handeler) DeleteCat_handler() fiber.Handler {
	return func (c *fiber.Ctx) error {
		val := c.Locals("userID")
		userID , ok := val.(uint)
		if !ok { 
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message" : "Invalid user ID",
			})
		}
    id_cat , err := strconv.Atoi(c.Params("id"))

		if err != nil { 
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message" : "Invalid category ID format",
				"error": err.Error(),
			})
		}

		var cat_detail domain.Cat_Delete;
		cat := &cat_detail
		cat.ID_User = userID
		cat.ID_Category = id_cat
		
		err = service.DeleteCat_service(cat , h.DB )	
		
		if err != nil {
  			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
     		"message" : "Failed to delete category",
				"error": err.Error(),
			})
		}
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message" : "Category deleted successfully",
		})
	}
}

func(h * CategoryOBG_handeler) PostCat_hanler() fiber.Handler { 
	return  func(c *fiber.Ctx) error {
		val := c.Locals("userID")
		userID , ok := val.(uint)

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message" : "Invalid user ID",
			})
		}
		var catPost domain.Cat_Action
		cat := &catPost
		err := c.BodyParser(cat)
		
		if err != nil { 
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message" : "Invalid request body",
				"error": err.Error(),
			})
		}
		cat.ID_User = userID;

		err = service.PostCat_service(cat ,h.DB)	
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Failed to create category",
				"error": err.Error(),
			})	
		}
		
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message" : "Category created successfully",
		})
	}
}


func (h *CategoryOBG_handeler) UpdateCat_handler() fiber.Handler {
	return  func(c *fiber.Ctx) error { 
	  val := c.Locals("userID")
		userID , ok := val.(uint)
		
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid user ID",
			})
		}
	  
		idCat , err := strconv.Atoi(c.Params("id"))
    
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid category ID format",
				"error": err.Error(),
			})
		}

		var cat_obj domain.Categories
		cat := &cat_obj
		
		if err := c.BodyParser(cat) ; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
				"error": err.Error(),
			})
		}

		cat.ID_User = int(userID)
    cat.ID_Category = idCat

		if err := service.UpdateCat_service(cat , h.DB) ; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update category",
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Category updated successfully",
		})
	}
}
