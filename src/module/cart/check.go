package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/model"
)

// CartCheckRequest struct
type CartCheckRequest struct {
	UserID int64 `json:"userID"`
}

// CartCheckResponse struct
type CartCheckResponse struct {
	Status string `json:"status"`
}

func CartCheck(cartCheckRequest *CartCheckRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	// get the cart by userID
	cart := new([]model.Cart)
	if err := tx.Where("user_id = ?", cartCheckRequest.UserID).Find(cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Add games to a inventory
	inventory := new([]model.Inventory)
	for i := 0; i < len(*cart); i++ {
		var item model.Inventory
		item.GameID = (*cart)[i].GameID
		item.UserID = (*cart)[i].UserID
		item.InventoryDateAdded = time.Now().String()[0:20]

		*inventory = append(*inventory, item)
	}

	if err := tx.Create(inventory).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete games from the cart
	if err := tx.Where("user_id = ?", cartCheckRequest.UserID).Delete(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func CartCheckOutput(w http.ResponseWriter, cartCheckResponse *CartCheckResponse) {
	jsonbyte, err := json.Marshal(cartCheckResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
