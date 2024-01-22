package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// UserRegistResponse struct
type GameUploadResponse struct {
	Status string `json:"status"`
}

// UserRegistRequest struct
type GameUploadRequest struct {
	UserID       int64  `json:"userID"`
	GamePrice    int64  `json:"gamePrice"`
	GameName     string `json:"gameName"`
	GameType     string `json:"gameType"`
	GameInfo     string `json:"gameInfo"`
	GameImg      string `json:"GameImg"`
	GameUploader string `json:"GameUploader"`
}

func GameUpload(gameUploadRequest *GameUploadRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	// get user's name
	user := new(model.User)
	if err := tx.Select("user_name").First(&user, gameUploadRequest.UserID).Error; err != nil {
		return err
	}

	// upload a new game
	game := new(model.Game)
	game.GameUploader = user.UserName
	game.GameImg = gameUploadRequest.GameImg
	game.GameInfo = gameUploadRequest.GameInfo
	game.GameName = gameUploadRequest.GameName
	game.GamePrice = gameUploadRequest.GamePrice
	game.GameType = gameUploadRequest.GameType

	if err := tx.Create(&game).Error; err != nil {
		tx.Rollback()
		return err
	}

	tag := new(model.Tag)

	if err := tx.First(&game, "game_name = ?", game.GameName).Error; err != nil {
		tx.Rollback()
		return err
	}
	tag.GameID = game.GameID
	tag.GameName = game.GameName

	// Add games to database by types
	for i := 0; i < len(game.GameType); i++ {
		if game.GameType[i] == ';' {
			// fmt.Println(tag.TagName)

			ID := new(model.Tag_list)
			if err := tx.Take(&ID, "tag_name = ?", tag.TagName).Error; err != nil {
				tx.Rollback()
				return err
			}

			tag.TagID = ID.TagID

			if err := tx.Create(&tag).Error; err != nil {
				tx.Rollback()
				return err
			}

			tag.TagName = ""
		} else {
			tag.TagName += string(game.GameType[i])
		}
	}

	tx.Commit()
	return tx.Error
}

func GameUploadOutput(w http.ResponseWriter, gameUploadResponse *GameUploadResponse) {
	jsonbyte, err := json.Marshal(gameUploadResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
