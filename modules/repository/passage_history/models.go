package passage_history

type PassageHistoryModel struct {
	ID          int `gorm:"primary_key"`
	UserID      int
	PassageID   int
	GameID      string
	Correctness bool

	User    userModel
	Passage PassageModel
}

func (PassageHistoryModel) TableName() string {
	return "passage_history"
}

type userModel struct {
	ID             int
	Email          *string
	DisplayName    *string
	Active         bool
	Image          *string
	PrivateProfile bool
}

func (userModel) TableName() string {
	return "user"
}

type PassageModel struct {
	ID    int
	Title string
}

func (PassageModel) TableName() string {
	return "passage"
}
