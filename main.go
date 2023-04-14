package main

import (
	"DTS/Chapter-3/final-project-myGram/repo"
	"DTS/Chapter-3/final-project-myGram/routers"
)

// @title My Gram API
// @version 1.0
// @description This is a final project Hacktiv8
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header -> Bearer
// @name Authorization
func main() {

	repo.StartDB()
	routers.StartServer().Run(":8080")
}
