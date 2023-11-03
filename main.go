package main

import (
	"github.com/gin-gonic/gin"

	"cp23kk1/common/databases"
	"cp23kk1/common/routes"
	"cp23kk1/modules/gameplays"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	// users.AutoMigrate()
	gameplays.AutoMigrate()
}

func main() {

	db := databases.Init()
	Migrate(db)

	router := gin.Default()
	
	routes.Run(router)
	
}
