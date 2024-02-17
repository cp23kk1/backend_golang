package score

import (
	"cp23kk1/common/databases"
	"time"

	"github.com/gin-gonic/gin"
)

type ScoreSerealizer struct {
	C *gin.Context
	databases.ScoreBoardModel
}
type ScoresSerealizer struct {
	C           *gin.Context
	scoreBoards []databases.ScoreBoardModel
}

type ScoreBoardResponse struct {
	Score       int       `json:"score"`
	Week        int       `json:"week"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Email       *string   `json:"email"`
	DisplayName *string   `json:"displayName"`
}

func ConvertToScoreBoardResponse(scoreBoard databases.ScoreBoardModel) ScoreBoardResponse {

	return ScoreBoardResponse{
		Score:       scoreBoard.Score,
		Week:        scoreBoard.Week,
		StartDate:   scoreBoard.StartDate,
		EndDate:     scoreBoard.EndDate,
		Email:       scoreBoard.User.Email,
		DisplayName: scoreBoard.User.DisplayName,
	}
}
func (self *ScoreSerealizer) Response() ScoreBoardResponse {
	scoreBoardModel := self.ScoreBoardModel

	scoreBoards := ConvertToScoreBoardResponse(scoreBoardModel)
	// if err != nil {
	// 	fmt.Println("Error mapping data:", err)
	// 	return scoreBoards
	// }
	return scoreBoards
}

func (self *ScoresSerealizer) Response() []ScoreBoardResponse {
	scoreBoardModel := self.scoreBoards
	var scoreBoards []ScoreBoardResponse

	for _, vocab := range scoreBoardModel {
		serializer := ScoreSerealizer{self.C, vocab}
		scoreBoards = append(scoreBoards, serializer.Response())

	}
	return scoreBoards
}
