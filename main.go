package main

import (
	"github.com/gin-gonic/gin"

	"cp23kk1/common/databases"
	"cp23kk1/common/routes"
	"cp23kk1/modules/repository/passage"
	passageHistoryRepo "cp23kk1/modules/repository/passage_history"
	scoreBoardRepo "cp23kk1/modules/repository/score_board"
	"cp23kk1/modules/repository/sentence"
	"cp23kk1/modules/repository/sentence_history"
	userRepo "cp23kk1/modules/repository/user"
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
	"cp23kk1/modules/repository/vocabulary_history"
	"cp23kk1/modules/repository/vocabulary_related"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	db.AutoMigrate(&userRepo.UserModel{})
	db.AutoMigrate(&scoreBoardRepo.ScoreBoardModel{})
	db.AutoMigrate(&vocabularyRepo.VocabularyModel{})
	db.AutoMigrate(&passage.PassageModel{})
	db.AutoMigrate(&sentence.SentenceModel{})
	db.AutoMigrate(&passageHistoryRepo.PassageHistoryModel{})
	db.AutoMigrate(&vocabulary_history.VocabularyHistoryModel{})
	db.AutoMigrate(&sentence_history.SentenceHistoryModel{})
	db.AutoMigrate(&vocabulary_related.VocabularyRelatedModel{})

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
