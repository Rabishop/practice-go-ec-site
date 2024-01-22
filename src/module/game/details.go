package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// GameIndexResponse struct
type GameDetailsResponse struct {
	Status    string     `json:"status"`
	GameItem  model.Game `json:"gameItem"`
	Inventory bool       `json:"inventory"`
}

// GameIndexRequest struct
type GameDetailsRequest struct {
	UserID   int64  `json:"userID"`
	GameName string `json:"gameName"`
}

func GameDetails(gameDetailsRequest *GameDetailsRequest, gameDetailsResponse *GameDetailsResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	// get games by gameName
	if err := tx.Where("game_name = ?", gameDetailsRequest.GameName).Take(&gameDetailsResponse.GameItem).Error; err != nil {
		tx.Rollback()
		return err
	}

	var count int64
	if err := tx.Where("user_id = ? AND game_id = ?", gameDetailsRequest.UserID, gameDetailsResponse.GameItem.GameID).Find(&model.Inventory{}).Count(&count).Error; err != nil {
		tx.Rollback()
		return err
	}

	if count > 0 {
		gameDetailsResponse.Inventory = true
	}

	tx.Commit()
	return tx.Error
}

func GameDetailsOutput(w http.ResponseWriter, gameDetailsResponse *GameDetailsResponse) {
	jsonbyte, err := json.Marshal(gameDetailsResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
