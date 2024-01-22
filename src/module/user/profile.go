package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

type UserProfileRequest struct {
	UserID int64 `json:"userID"`
}

type UserProfileResponse struct {
	Status        string `json:"status"`
	UserName      string `json:"userName"`
	UserGameCount int    `json:"userGameCount"`
	UserPortrait  string `json:"userPortrait"`
}

type UserUploadPortraitRequest struct {
	UserID       int64  `json:"userID"`
	UserPortrait string `json:"userPortrait"`
}

type UserUploadPortraitResponse struct {
	Status string `json:"status"`
}

func Profile(userProfileRequest *UserProfileRequest, userProfileResponse *UserProfileResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User

	// get user's profile
	if err := tx.Where("user_id = ?", userProfileRequest.UserID).Take(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	userProfileResponse.UserName = user.UserName
	userProfileResponse.UserPortrait = user.UserPortrait
	userProfileResponse.UserGameCount = 0

	tx.Commit()
	return tx.Error
}

func ProfileOutput(w http.ResponseWriter, userProflieResponse *UserProfileResponse) {
	jsonbyte, err := json.Marshal(userProflieResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}

func UploadPortrait(userUploadPortraitRequest *UserUploadPortraitRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	// get user's profile
	if err := tx.Table("user").Where("user_id = ?", userUploadPortraitRequest.UserID).Update("user_portrait", userUploadPortraitRequest.UserPortrait).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func UploadPortraitOutput(w http.ResponseWriter, userUploadPortraitResponse *UserUploadPortraitResponse) {
	jsonbyte, err := json.Marshal(userUploadPortraitResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
