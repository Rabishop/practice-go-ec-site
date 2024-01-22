package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// CartRemoveRequest struct
type CartRemoveRequest struct {
	UserID int64 `json:"userID"`
	GameID int64 `json:"gameID"`
}

// CartRemoveResponse struct
type CartRemoveResponse struct {
	Status string `json:"status"`
}

func CartRemove(cartRemoveRequest *CartRemoveRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	cart := new(model.Cart)
	cart.UserID = cartRemoveRequest.UserID
	cart.GameID = cartRemoveRequest.GameID

	// Delete games
	if err := tx.Where("user_id = ? AND game_id = ?", cart.UserID, cart.GameID).Delete(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func CartRemoveOutput(w http.ResponseWriter, cartRemoveResponse *CartRemoveResponse) {
	jsonbyte, err := json.Marshal(cartRemoveResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
