package middleware

import (
	"fmt"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifySession() fiber.Handler { 
	return func(c *fiber.Ctx) error { 

   user_cookie	 := c.Cookies("jwt_Token")
	 if user_cookie == "" { 
  		return c.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"messege" : "Unauthorized: Missing token cookie",
			})
	 }
		
	user_claims := new(domain.User_Claim)


	token, err := jwt.ParseWithClaims(user_cookie , user_claims , func(token *jwt.Token) ( any , error) {
		if _ , ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil , fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		 return []byte("SECRET_KEY") , nil ;
	 })
	 if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"messege" : "Unauthorized: Invalid or expired token",
			})
	 }
	userID := user_claims.ID

	c.Locals("userID", userID);

	 return c.Next()
	}
}
