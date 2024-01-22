package model

type User struct {
	UserID       int64  `gorm:"column:user_id;AUTO_INCREMENT;" json:"userID"`
	UserAccount  string `gorm:"column:user_account;" json:"userAccount"`
	UserPassword string `gorm:"column:user_password;" json:"userPassword"`
	UserName     string `gorm:"column:user_name;" json:"userName"`
	UserPortrait string `gorm:"column:user_portrait;" json:"userPortrait"`
}

type UserID struct {
	UserID int64 `gorm:"column:user_id;AUTO_INCREMENT;" json:"userID"`
}

// Return TableName
func (User) TableName() string {
	return "user"
}
