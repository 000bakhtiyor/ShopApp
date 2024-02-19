package api

import (
	"shopapp/handlers"
	"shopapp/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func ApiRoutes(h handlers.HandlerInterface) {
	engine := gin.Default()

	engine.Use(cors.Default())
	
	engine.POST("/register", h.Register) // admin and user
	engine.POST("/login", h.Login) // admin and user
	engine.GET("/log-out", middlewares.Login, h.LogOut) // admin and user
	engine.POST("/create-customer", middlewares.Login, h.CreateCustomer) // user

	engine.POST("/create-category", middlewares.Login, h.CreateCategory) // admin 
	engine.DELETE("/delete-category", middlewares.Login, h.DeleteCategory) // admin 

	engine.POST("/create-product", middlewares.Login, h.CreateProduct) // user
	engine.DELETE("/delete-product", middlewares.Login, h.DeleteProduct) // user
	engine.GET("/products", h.Products) // all
	engine.POST("/add-image-to-product", h.AddImageToProduct) // user

	engine.POST("/create-order", middlewares.Login, h.CreateOrder) // user
	engine.POST("/create-order-item", middlewares.Login, h.CreateOrderItem) // user
	engine.POST("/create-payment-method", middlewares.Login, h.CreatePaymentMethod) // user

	engine.POST("/create-review", middlewares.Login, h.CreateReview) // user

	engine.Run()
}