package handlers

import (
	"net/http"
	"shopapp/config"
	"shopapp/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h HandlerStruct) DeleteProduct(ctx *gin.Context)  {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.db.DeleteProduct(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product is deleted",
	})
}

func (h HandlerStruct) AddImageToProduct(ctx *gin.Context)  {
	var newImage models.Image
	err := ctx.ShouldBindJSON(&newImage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.db.AddImageToProduct(newImage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} 

	ctx.JSON(200, gin.H{
		"message": "Image is added",
	})
}

func (h HandlerStruct) CreateReview(ctx *gin.Context)  {
	var newReview models.Review
	err := ctx.ShouldBindJSON(&newReview)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	newReview.Timestamp = time.Now()
	tokenString, err := ctx.Cookie("Authorization")
	if tokenString == "" || err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not login" + err.Error(),
		})
		return
	}

	claims, err := config.GetClaimsFromToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not login:"+err.Error(),
		})
		return
	}

	id := claims["user_id"]

	newReview.CustomerID = int(id.(float64))
	err = h.db.CreateReveiw(newReview)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Review is added",
	})
}

func (h HandlerStruct) CreatePaymentMethod(ctx *gin.Context)  {
	var newPaymentMethod models.PaymentMethod
	err := ctx.ShouldBindJSON(&newPaymentMethod)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// This one is error way
	newPaymentMethod.ExpirationDate = time.Now()
	tokenString, err := ctx.Cookie("Authorization")
	if tokenString == "" || err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not login" + err.Error(),
		})
		return
	}

	claims, err := config.GetClaimsFromToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not login:"+err.Error(),
		})
		return
	}

	id := claims["user_id"]

	newPaymentMethod.CustomerID = int(id.(float64))
	err = h.db.CreatePaymentMethod(newPaymentMethod)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Payment method is added",
	})
}

func (h HandlerStruct) CreateOrderItem(ctx *gin.Context)  {
	var newOrderItem models.OrderItem

	err := ctx.ShouldBindJSON(&newOrderItem)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.db.CreateOrderItem(newOrderItem)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":"Order item is added",
	})
}

func (h HandlerStruct) CreateOrder(ctx *gin.Context)  {
	var newOrder models.Order
	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	newOrder.OrderDate = time.Now()
	tokenString, err := ctx.Cookie("Authorization")
	if tokenString == "" || err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not login" + err.Error(),
		})
		return
	}

	claims, err := config.GetClaimsFromToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not login:"+err.Error(),
		})
		return
	}

	id := claims["user_id"]

	newOrder.CustomerID = int(id.(float64))
	err = h.db.CreateOrder(newOrder)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Order is added",
	})
}

func (h HandlerStruct) CreateCustomer(ctx *gin.Context){
	var newCustomer models.Customer

	err := ctx.ShouldBindJSON(&newCustomer)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	tokenString, err := ctx.Cookie("Authorization")
	if tokenString == "" || err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not login" + err.Error(),
		})
		return
	}

	claims, err := config.GetClaimsFromToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not login:"+err.Error(),
		})
		return
	}

	id := claims["user_id"]

	newCustomer.UserID = int(id.(float64))

	err = h.db.CreateCustomer(newCustomer)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": newCustomer.FirstName + " is added",
	})
}

func (h HandlerStruct) CreateProduct(ctx *gin.Context){
	var newProduct models.Product
	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.db.CreateProduct(newProduct)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": newProduct.Name + " added.",
	})
}

func (h HandlerStruct) CreateCategory(ctx *gin.Context){
	var categoryName models.Category
	err := ctx.ShouldBindJSON(&categoryName)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message":err.Error(),
		})
		return
	}

	err = h.db.CreateCategory(categoryName.Name)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":categoryName.Name + " added",
	})	
}

func (h HandlerStruct) DeleteCategory(ctx *gin.Context){
	var categoryName string
	categoryName, exists := ctx.GetQuery("category-name")
	if !exists {
		ctx.JSON(400, gin.H{
			"message":"value is empty string",
		})
		return
	}

	err := h.db.DeleteCategory(categoryName)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": categoryName + " deleted",
	})
}

func (h HandlerStruct) Register(ctx *gin.Context){
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message":err.Error(),
		})
		return 
	}	
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message":"Error to hashing",
		})
		return 
	}
	user.SetPassword(string(hash))

	err = h.db.UserRegister(user)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":"User has been registered.",
	})
}

func (h HandlerStruct) Login(ctx *gin.Context){

	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message":err.Error(),
		})
		return 
	}	

	newUser, err := h.db.UserLogin(user)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message":err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(user.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword{
			ctx.JSON(401, gin.H{
				"message":"password is incorrect",
			})
			return
		}
	}

	tokenString, err := config.CreateToken(newUser)
	if err != nil {
		ctx.JSON(401, gin.H{
			"message":err.Error(),
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30,"", "", false, true)

	ctx.JSON(200, gin.H{})
}

func (h HandlerStruct) LogOut(ctx *gin.Context){
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
}

func (h HandlerStruct) Products(ctx *gin.Context){
	var productList []models.Product

	productList, err := h.db.Products()
	if err != nil {
		ctx.JSON(400, gin.H{
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(200, productList)
}