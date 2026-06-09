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
// GetallUser_handler
// @Summary      Get all users
// @Description  Retrieve a list of all users containing only their IDs and Emails for testing purposes.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200      {object}  map[string]interface{} "Success with user list"
// @Failure      400      {string}  string                 "Bad Request"
// @Router       /api/v1/user/ [get]
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

// RegisterUser_handler
// @Summary      Register a new user
// @Description  Create a new user account using an email and password.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request  body      domain.User             true  "Registration credentials"
// @Success      202      {object}  map[string]interface{}  "Successfully Registered"
// @Failure      400      {string}  string                  "Bad Request / Validation Error"
// @Router       /api/v1/user/register [post]
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

// LoginUser_handler
// @Summary      User Login
// @Description  Authenticate user with email and password, returning a HttpOnly JWT Token via Cookie.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request  body      domain.User_Login       true  "Login credentials"
// @Success      202      {object}  map[string]interface{}  "Login Successful (Sets jwt_Token cookie)"
// @Failure      400      {string}  string                  "Invalid credentials or bad payload"
// @Router       /api/v1/user/login [post]
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

// DeleteUser_handler
// @Summary      Delete a user by ID
// @Description  Remove a user account permanently from the database using their ID.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id       path      int                     true  "User ID to delete"
// @Success      202      {object}  map[string]interface{}  "Successfully Deleted"
// @Failure      400      {string}  string                  "Invalid ID or user not found"
// @Router       /api/v1/user/{id} [delete]
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
