package database

import (
	"database/sql"
	"errors"
	"shopapp/models"
)

func (db DBStruct) DeleteProduct(product models.Product) error {
	query := `DELETE FROM products WHERE product_id = $1 OR name=$2`

	_, err := db.sqlDB.Exec(query, product.ProductID, product.Name)
	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) AddImageToProduct(image models.Image) error {
	query := `INSERT INTO images(
		product_id, url 
	) VALUES (
		$1, $2
	)`

	_, err := db.sqlDB.Exec(query, image.ProductID, image.URL)
	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) CreateReveiw(review models.Review) error {
	
	query := `INSERT INTO reviews(
		product_id, customer_id, rating, comment, timestamp 
	) VALUES (
		$1, $2, $3, $4, $5
	)`

	_, err := db.sqlDB.Exec(query, 
		review.ProductID,
		review.CustomerID,
		review.Rating,
		review.Comment,
		review.Timestamp,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) CreatePaymentMethod(newPaymentMethod models.PaymentMethod) error {
	
	query := `INSERT INTO payment_methods (
		customer_id, card_number, expiration_date, cvv 
	) VALUES (
		$1, $2, $3, $4
	)`

	_, err := db.sqlDB.Exec(query, 
		newPaymentMethod.CustomerID, 
		newPaymentMethod.CardNumber, 
		newPaymentMethod.ExpirationDate,
		newPaymentMethod.CVV,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) CreateOrderItem(newOrderItem models.OrderItem) error {
	
	query := `INSERT INTO order_items (
		order_id, product_id, quantity, subtotal
	) VALUES (
		$1, $2, $3, $4
	)`

	_, err := db.sqlDB.Exec(query, 
		newOrderItem.OrderID, 
		newOrderItem.ProductID, 
		newOrderItem.Quantity, 
		newOrderItem.Subtotal)
	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) CreateOrder(newOrder models.Order) error {
	query := `INSERT INTO orders(
		customer_id, order_date, status 
	) VALUES (
		$1, $2, $3
	)`

	_, err := db.sqlDB.Exec(query, newOrder.CustomerID, newOrder.OrderDate, newOrder.Status)
	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) CreateProduct(newProduct models.Product) error {
	query := `INSERT INTO products(
		name, description, price, stock_quantity, category_id
	) VALUES (
		$1,$2,$3,$4,$5
	)`

	_, err := db.sqlDB.Exec(query, 
		newProduct.Name,
		newProduct.Description,
		newProduct.Price,
		newProduct.StockQuantity,
		newProduct.CategoryID,
		)	

	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) Products() ([]models.Product, error) {
	query := "SELECT * FROM products"
	rows, err := db.sqlDB.Query(query)
	if err != nil {
		return []models.Product{}, err
	}
	var product models.Product
	var productList []models.Product
	for rows.Next() {
		if err = rows.Scan(&product.CategoryID,
							&product.Name,
							&product.Description,
							&product.Price,
							&product.StockQuantity,
							&product.CategoryID); err != nil{
			return []models.Product{}, err
		}
		productList = append(productList, product)
	}

	return productList, nil
}

func (db DBStruct) CreateCategory(categoryName string) error {
	query := "INSERT INTO Categories (name) VALUES ($1)"

	_, err := db.sqlDB.Exec(query, categoryName)
	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) DeleteCategory(categoryName string) error {
	
	query := "DELETE FROM categories WHERE name = $1"

	_, err := db.sqlDB.Exec(query, categoryName)
	if err != nil {
		return err
	} 

	return nil
}

func (db DBStruct) UserRegister(newUser models.User) error {
	_, err := db.sqlDB.Exec(
		`INSERT INTO users (
			username, password
		)
		VALUES (
			$1,$2
		)`,
		newUser.Username, newUser.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db DBStruct) UserLogin(User models.User) (models.User, error){

	var getUser models.User

	sqlRow := db.sqlDB.QueryRow(
		`SELECT * FROM users WHERE username=$1`,
		User.Username,
	)
	if err:=sqlRow.Scan(&getUser.UserID, &getUser.Username, &getUser.Password); err!=nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return getUser, nil
}

func (db DBStruct) CreateCustomer(newCustomer models.Customer) error {

	query := `INSERT INTO customers (
		first_name, last_name, email, phone_number, shipping_address, billing_address, user_id
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7
	)`

	_, err := db.sqlDB.Exec(query,
		newCustomer.FirstName,
		newCustomer.LastName,
		newCustomer.Email,
		newCustomer.PhoneNumber,
		newCustomer.ShippingAddress,
		newCustomer.BillingAddress,
		newCustomer.UserID,
		)	
	if err != nil {
		return err
	}

	return nil 
}