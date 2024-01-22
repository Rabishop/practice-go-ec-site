package model

type Tag struct {
	TagID    int64  `gorm:"column:tag_id;AUTO_INCREMENT;" json:"tagID"`
	GameID   int64  `gorm:"column:game_id;" json:"gameID"`
	GameName string `gorm:"column:game_name;" json:"gameName"`
	TagName  string `gorm:"column:tag_name;" json:"tagName"`
}

// Return TableName
func (Tag) TableName() string {
	return "tag"
}
