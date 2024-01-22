package model

type Cart struct {
	UserID        int64  `gorm:"column:user_id;" json:"userID"`
	GameID        int64  `gorm:"column:game_id;" json:"gameID"`
	CartDateAdded string `gorm:"column:cart_date_added;" json:"cartDateAdded"`
}

// Return TableName
func (Cart) TableName() string {
	return "cart"
}
