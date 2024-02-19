package models

import (
 "time"
)

// Categories table
type Category struct {
 CategoryID int    `json:"category_id"`
 Name       string `json:"name"`
}

// Products table
type Product struct {
 ProductID      int     `json:"product_id"`
 Name           string  `json:"name"`
 Description    string  `json:"description"`
 Price          float64 `json:"price"`
 StockQuantity  int     `json:"stock_quantity"`
 CategoryID     int     `json:"category_id"`
}

// Customers table
type Customer struct {
 CustomerID       int    `json:"customer_id"`
 FirstName        string `json:"first_name"`
 LastName         string `json:"last_name"`
 Email            string `json:"email"`
 PhoneNumber      string `json:"phone_number"`
 ShippingAddress  string `json:"shipping_address"`
 BillingAddress   string `json:"billing_address"`
 UserID			  int    `json:"user_id"`
}

// Users table
type User struct {
 UserID     int    `json:"user_id"`
 Username   string `json:"username"`
 Password   string `json:"password"`
}

// Orders table
type Order struct {
 OrderID     int       `json:"order_id"`
 CustomerID  int       `json:"customer_id"`
 OrderDate   time.Time `json:"order_date"`
 Status      string    `json:"status"`
}

// Order_Items table
type OrderItem struct {
 OrderItemID int     `json:"order_item_id"`
 OrderID     int     `json:"order_id"`
 ProductID   int     `json:"product_id"`
 Quantity    int     `json:"quantity"`
 Subtotal    float64 `json:"subtotal"`
}

// Payment_Methods table
type PaymentMethod struct {
 PaymentMethodID int    `json:"payment_method_id"`
 CustomerID      int    `json:"customer_id"`
 CardNumber      string `json:"card_number"`
 ExpirationDate  time.Time `json:"expiration_date"`
 CVV             string `json:"cvv"`
}

// Reviews table
type Review struct {
 ReviewID   int       `json:"review_id"`
 ProductID  int       `json:"product_id"`
 CustomerID int       `json:"customer_id"`
 Rating     int       `json:"rating"`
 Comment    string    `json:"comment"`
 Timestamp  time.Time `json:"timestamp"`
}

// Images table
type Image struct {
 ImageID   int    `json:"image_id"`
 ProductID int    `json:"product_id"`
 URL       string `json:"url"`
}