package handlers

import (
	"shopapp/database"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	Register (*gin.Context)
	Login (*gin.Context)
	LogOut (*gin.Context)
	CreateCustomer (*gin.Context)

	CreateCategory (*gin.Context)
	DeleteCategory (*gin.Context)

	CreateOrder (*gin.Context)
	CreateOrderItem (*gin.Context)
	CreatePaymentMethod (*gin.Context)

	CreateReview (*gin.Context)

	CreateProduct (*gin.Context)
	DeleteProduct (*gin.Context)
	Products (*gin.Context)
	AddImageToProduct (*gin.Context)
}

type HandlerStruct struct {
	db database.DBInterface
}

func CreateHandler(db database.DBInterface) (HandlerInterface){
	return HandlerStruct{
		db: db,
	}
}