package passage_history

type PassageHistoryModel struct {
	ID          int    `gorm:"primary_key"`
	UserID      int    `gorm:"not null"`
	PassageID   int    `gorm:"not null"`
	GameID      string `gorm:"not null;column:game_id;size:45"`
	Correctness bool   `gorm:"not null"`
}
