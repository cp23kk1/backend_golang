package main

import (
	"github.com/gin-gonic/gin"

	"cp23kk1/common/databases"
	"cp23kk1/common/routes"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	db.AutoMigrate(&databases.UserModel{})
	db.AutoMigrate(&databases.DifficultyModel{})
	db.AutoMigrate(&databases.ScoreBoardModel{})
	db.AutoMigrate(&databases.VocabularyModel{})
	db.AutoMigrate(&databases.PassageModel{})
	db.AutoMigrate(&databases.SentenceModel{})
	db.AutoMigrate(&databases.PassageHistoryModel{})
	db.AutoMigrate(&databases.VocabularyHistoryModel{})
	db.AutoMigrate(&databases.SentenceHistoryModel{})
	db.AutoMigrate(&databases.VocabularyRelatedModel{})

	// db.AutoMigrate(&userRepo.UserModel{},
	// 	&scoreBoardRepo.ScoreBoardModel{},
	// 	&vocabularyRepo.VocabularyModel{},
	// 	&passage.PassageModel{},
	// 	&sentence.SentenceModel{},
	// 	&passageHistoryRepo.PassageHistoryModel{},
	// 	&vocabulary_history.VocabularyHistoryModel{},
	// 	&sentence_history.SentenceHistoryModel{},
	// 	&vocabulary_related.VocabularyRelatedModel{},
	// )
	// userRepo.AutoMigrate(db)
	// scoreBoardRepo.AutoMigrate(db)
	// vocabularyRepo.AutoMigrate(db)
	// passage.AutoMigrate(db)
	// passageHistoryRepo.AutoMigrate(db)
	// vocabulary_history.AutoMigrate(db)
}

func main() {

	db := databases.Init()
	Migrate(db)

	router := gin.Default()

	routes.Run(router)

}
