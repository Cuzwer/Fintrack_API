package handler

import (
	"strconv"
	"time"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserHadler struct {
	 DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHadler{
	var user_object UserHadler
	u := &user_object
	u.DB = db

	return u
}


func GetallUser_handler(db *gorm.DB) fiber.Handler{
	return func(c *fiber.Ctx) error { 
   var user_object []domain.User_List
	 user := &user_object

	 if err := service.GetAlluser_service(user,db) ; err != nil { 
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

   return c.Status(fiber.StatusOK).JSON(fiber.Map{
   "message" : fiber.StatusOK,
	 "value" : user,
	 })
	}
}


func (h *UserHadler) RegisterUser_handler()  fiber.Handler {
	return  func (c *fiber.Ctx) error {
		
 		var user_object domain.User;
		user := &user_object;
    
		if err := c.BodyParser(user) ; err != nil { 
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		
		err :=service.RegisterUser_service(user,h.DB)
		if err != nil { 
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "Succes fully Registurtion",
			"Status" : fiber.StatusAccepted,
		})
}
}

func (h *UserHadler) LoginUser_handler() fiber.Handler { 
	return func(c *fiber.Ctx) error { 
		user := new(domain.User)
		
		if err := c.BodyParser(user) ;  err != nil { 
			return  c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		
		if err := service.LoginUser_Service(user, h.DB) ; err != nil { 
			return  c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		jwt_Token := jwt.NewWithClaims(jwt.SigningMethodHS256  , jwt.MapClaims {
			"user_id" : user.ID,
			"exp" : jwt.NewNumericDate(time.Now().Add(time.Hour * 24 )),
		})
		tokenString , _ := jwt_Token.SignedString([]byte("SECRET_KEY"))
	 
		c.Cookie(&fiber.Cookie{
			Name : "jwt_Token",
			Value: tokenString,
			Expires: time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure: false,
			SameSite: "Lax",
		})
    return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "Login In Successfully",
			"status" : fiber.StatusAccepted,
		})
	}  
}

func (h *UserHadler) DeleteUser_handler() fiber.Handler { 
	return  func(c *fiber.Ctx) error { 
		user_id , err := strconv.Atoi(c.Params("id"));
		
		if err != nil { 
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())	
		}
  	
		err = service.DeleteUserById_service(user_id, h.DB)
		
		if err != nil { 
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())	
		}
			
		return  c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "Successfully Delete user",
		})
	}
}
