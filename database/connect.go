package database

import (
	"database/sql"
	"fmt"
	"shopapp/config"
	"shopapp/models"

	_ "github.com/lib/pq"
)

type DBInterface interface {
	UserRegister (models.User) error
	UserLogin (models.User) (models.User, error)
	CreateCustomer (models.Customer) error

	CreateProduct (models.Product) error
	DeleteProduct (models.Product) error
	Products ()([]models.Product, error)
	AddImageToProduct (models.Image) error

	CreateOrder (models.Order) error
	CreateOrderItem (models.OrderItem) error
	CreatePaymentMethod (models.PaymentMethod) error

	CreateReveiw (models.Review) error

	CreateCategory (string) error
	DeleteCategory (string) error
}

type DBStruct struct {
	sqlDB *sql.DB
}

func CreateDB() (DBInterface){
	return DBStruct{
		sqlDB: ConnectToPostgres(),
	}
}

func ConnectToPostgres() (*sql.DB) {

	cfg := config.CreateConfig()

	openDB, err := sql.Open(cfg.DriverName, cfg.DataSourceName.Get())
	if err != nil {
		fmt.Println("Ulanishda xatolik: ", err)
		return nil
	}

	// defer openDB.Close()

	err = openDB.Ping()
	if err != nil {
		fmt.Println("Pingda xatolik: ", err)
		return nil
	}
  
	fmt.Println("Successfully connected!")

	return openDB
}