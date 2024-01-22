package model

type Game struct {
	GameID       int64  `gorm:"column:game_id;AUTO_INCREMENT;" json:"gameID"`
	GameStar     int64  `gorm:"column:game_star;" json:"gameStar"`
	GamePrice    int64  `gorm:"column:game_price;" json:"gamePrice"`
	GameName     string `gorm:"column:game_name;" json:"gameName"`
	GameType     string `gorm:"column:game_type;" json:"gameType"`
	GameInfo     string `gorm:"column:game_info;" json:"gameInfo"`
	GameImg      string `gorm:"column:game_img;" json:"gameImg"`
	GameUploader string `gorm:"column:game_uploader;" json:"gameUploader"`
}

// Return TableName
func (Game) TableName() string {
	return "game"
}
