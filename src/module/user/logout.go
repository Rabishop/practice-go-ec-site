package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// UserLogoutResponse struct
type UserLogoutResponse struct {
	Status string `json:"status"`
}

// UserLogoutRequest struct
type UserLogoutRequest struct {
	UserID int64 `json:"userID"`
}

// Delete the cookies to log out
func Logout(userLogoutRequest *UserLogoutRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	// user login
	if err := tx.Delete(&model.Session{}, "user_id = ?", userLogoutRequest.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func LogoutOutput(w http.ResponseWriter, userLoginResponse *UserLoginResponse) {
	jsonbyte, err := json.Marshal(userLoginResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
