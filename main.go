package main

import (
	"DTS/Chapter-3/final-project-myGram/repo"
	"DTS/Chapter-3/final-project-myGram/routers"
)

func main() {

	repo.StartDB()
	routers.StartServer().Run(":8080")
}
