package main

import (
	"github.com/gin-gonic/gin"

	"cp23kk1/common/databases"
	"cp23kk1/common/routes"
	"cp23kk1/modules/repository/passage"
	passageHistoryRepo "cp23kk1/modules/repository/passage_history"
	scoreBoardRepo "cp23kk1/modules/repository/score_board"
	userRepo "cp23kk1/modules/repository/user"
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
	"cp23kk1/modules/repository/vocabulary_history"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	userRepo.AutoMigrate(db)
	scoreBoardRepo.AutoMigrate(db)
	vocabularyRepo.AutoMigrate(db)
	passage.AutoMigrate(db)
	passageHistoryRepo.AutoMigrate(db)
	vocabulary_history.AutoMigrate(db)
}

func main() {

	db := databases.Init()
	Migrate(db)

	router := gin.Default()
	routes.Run(router)

}
