package model

type Session struct {
	UserID          int64  `gorm:"column:user_id;" json:"userID"`
	SessionID       string `gorm:"column:session_id;" json:"sessionID"`
	SessionLastTime string `gorm:"column:session_last_time;" json:"sessionLastTime"`
}

// Return TableName
func (Session) TableName() string {
	return "session"
}
