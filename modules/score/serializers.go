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
	ID          uint      `json:"scoreId"`
	Score       int       `json:"score"`
	Week        int       `json:"week"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Email       *string   `json:"email"`
	DisplayName *string   `json:"displayName"`
}

type BestScoreResponse struct {
	Score int     `json:"score"`
	Mode  *string `json:"mode"`
}

func ConvertToScoreBoardResponse(scoreBoard databases.ScoreBoardModel) ScoreBoardResponse {

	return ScoreBoardResponse{
		ID:          scoreBoard.ID,
		Score:       scoreBoard.Score,
		Week:        scoreBoard.Week,
		StartDate:   scoreBoard.StartDate,
		EndDate:     scoreBoard.EndDate,
		Email:       scoreBoard.User.Email,
		DisplayName: scoreBoard.User.DisplayName,
	}
}

func ConvertToBestScoreResponse(scoreBoard databases.ScoreBoardModel) BestScoreResponse {

	return BestScoreResponse{
		Score: scoreBoard.Score,
		Mode:  nil,
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

	for _, score := range scoreBoardModel {
		serializer := ScoreSerealizer{self.C, score}
		scoreBoards = append(scoreBoards, serializer.Response())

	}
	return scoreBoards
}

func (self *ScoreSerealizer) BestScoreResponse() BestScoreResponse {
	scoreBoardModel := self.ScoreBoardModel

	scoreBoards := ConvertToBestScoreResponse(scoreBoardModel)
	// if err != nil {
	// 	fmt.Println("Error mapping data:", err)
	// 	return scoreBoards
	// }
	return scoreBoards
}

func (self *ScoresSerealizer) BestScoreResponse() []BestScoreResponse {
	scoreBoardModel := self.scoreBoards
	var scoreBoards []BestScoreResponse

	for _, score := range scoreBoardModel {
		serializer := ScoreSerealizer{self.C, score}
		scoreBoards = append(scoreBoards, serializer.BestScoreResponse())

	}
	return scoreBoards
}
