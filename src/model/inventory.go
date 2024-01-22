package model

type Inventory struct {
	UserID               int64  `gorm:"column:user_id;" json:"userID"`
	GameID               int64  `gorm:"column:game_id;" json:"gameID"`
	InventoryDateAdded   string `gorm:"column:inventory_date_added;" json:"InventoryDateAdded"`
	InventoryHoursPlayed int64  `gorm:"column:inventory_hours_played;" json:"InventoryHoursPlayed"`
}

// Return TableName
func (Inventory) TableName() string {
	return "inventory"
}
