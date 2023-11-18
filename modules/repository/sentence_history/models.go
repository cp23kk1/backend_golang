package sentence_history

type SentenceHistoryModel struct {
	ID          int `gorm:"primaryKey"`
	UserID      int
	SentenceID  int
	GameID      string
	Correctness bool

	User     userModel
	Sentence SentenceModel
}

func (SentenceHistoryModel) TableName() string {
	return "sentence_history"
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

type SentenceModel struct {
	ID        int
	PassageID int
	Sequence  int
	Text      string
	Meaning   string
}

func (SentenceModel) TableName() string {
	return "sentence"
}
