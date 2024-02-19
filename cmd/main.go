package main

import (
	"shopapp/api"
	"shopapp/database"
	"shopapp/handlers"
)

func main (){

	db := database.CreateDB()
	h := handlers.CreateHandler(db)
	api.ApiRoutes(h)

}