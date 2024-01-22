package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// GameIndexResponse struct
type GameBrowserResponse struct {
	Status   string       `json:"status"`
	GameItem []model.Game `json:"gameItem"`
}

// GameIndexRequest struct
type GameBrowserResqust struct {
	GameType string `json:"gameType"`
}

func GameBrowser(gameBrowserResqust *GameBrowserResqust, gameBrowserResponse *GameBrowserResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	// View games by types
	if err := tx.Where("game_type like ?", "%"+gameBrowserResqust.GameType+"%").Find(&gameBrowserResponse.GameItem).Error; err != nil {
		tx.Rollback()
		return err
	}

	// fmt.Println(gameBrowserResponse.GameItem)

	tx.Commit()
	return tx.Error
}

func GameBrowserOutput(w http.ResponseWriter, gameBrowserResponse *GameBrowserResponse) {
	jsonbyte, err := json.Marshal(gameBrowserResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
