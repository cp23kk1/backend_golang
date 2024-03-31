package users

import (
	"cp23kk1/common"
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/passage_history"
	"cp23kk1/modules/repository/sentence_history"
	userRepo "cp23kk1/modules/repository/user"
	"cp23kk1/modules/repository/vocabulary_history"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserUseCase struct{}

func getUser(id uint) (*databases.UserModel, error) {
	userRepository := userRepo.NewUserRepository(databases.GetDB())
	return userRepository.FindUserByID(id)
}
func getStatisticService(c *gin.Context, userId int) {
	fmt.Println("userId ", userId)
	db := databases.GetDB()
	passageHistoryRepository := passage_history.NewPassageHistoryRepository(db)
	vocabularyHistoryRepository := vocabulary_history.NewVocabularyHistoryRepository(db)
	sentenceHistoryRepository := sentence_history.NewSentenceHistoryRepository(db)

	passage, err := passageHistoryRepository.FindPassageHistoriesByUserID(userId)
	passageCorrect, err := passageHistoryRepository.FindPassageHistoriesByUserIDAndCorrect(userId)
	vocabulary, err := vocabularyHistoryRepository.FindVocabulariesByUserID(userId)
	vocabularyCorrect, err := vocabularyHistoryRepository.FindVocabulariesByUserIDAndCorrect(userId)
	sentence, err := sentenceHistoryRepository.FindSentenceHistoriesByUserID(userId)
	sentenceCorrect, err := sentenceHistoryRepository.FindSentenceHistoriesByUserIDAndCorrect(userId)

	countVocabulary := vocabularyHistoryRepository.FindCountVocabularyHistoryGroupByPOS(userId)
	countVocabularyCorrect := vocabularyHistoryRepository.FindCountVocabularyHistoryGroupByPOSAndCorrect(userId)
	countSentence := sentenceHistoryRepository.FindCountSentenceHistoryGroupByPOS(userId)
	countSentenceCorrect := sentenceHistoryRepository.FindCountSentenceHistoryGroupByPOSAndCorrect(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: err.Error(), Status: "error"}, map[string]interface{}{}))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get Statistic successfully", Status: "success"}, map[string]interface{}{"overall": (len(passage) + len(vocabulary) + len(sentence)),
		"overallCorrect":  (len(passageCorrect) + len(vocabularyCorrect) + len(sentenceCorrect)),
		"countVocabulary": countVocabulary, "countVocabularyCorrect": countVocabularyCorrect,
		"countSentence": countSentence, "countSentenceCorrect": countSentenceCorrect,
		"countPassage": len(passage), "countPassageCorrect": len(passageCorrect),
	}))
	return

}
