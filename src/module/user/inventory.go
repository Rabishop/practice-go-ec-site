package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// UserInventoryResponse struct
type UserInventoryResponse struct {
	Status   string            `json:"status"`
	GameList []model.Inventory `json:"gameList"`
	ItemList []model.Game      `json:"itemList"`
}

// UserInventoryRequest struct
type UserInventoryRequest struct {
	UserID int64 `json:"userID"`
}

func Inventory(userInventoryRequest *UserInventoryRequest, userInventoryResponse *UserInventoryResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	gameItem := new(model.Game)

	// get user's games
	if err := tx.Where("user_id = ?", userInventoryRequest.UserID).Find(&userInventoryResponse.GameList).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := 0; i < len(userInventoryResponse.GameList); i++ {
		if err := tx.Where("game_id = ?", userInventoryResponse.GameList[i].GameID).Take(&gameItem).Error; err != nil {
			tx.Rollback()
			return err
		}
		// fmt.Println(gameItem)
		userInventoryResponse.ItemList = append(userInventoryResponse.ItemList, *gameItem)
	}

	// fmt.Println(userInventoryResponse.GameList)

	tx.Commit()
	return tx.Error
}

func InventoryOutput(w http.ResponseWriter, userInventoryResponse *UserInventoryResponse) {
	jsonbyte, err := json.Marshal(userInventoryResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
