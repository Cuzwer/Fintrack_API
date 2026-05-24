package middleware

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func NewCorsMiddleware () fiber.Handler { 

	newCors := cors.New(cors.Config{
    AllowOrigins: "http://127.0.0.1:5000",
    AllowCredentials: true,      
	})
	return  newCors
}
