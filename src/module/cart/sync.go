package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// CartUploadRequest struct
type CartSyncRequest struct {
	UserID        int64    `json:"userID"`
	GameID        []int64  `json:"gameID"`
	CartDateAdded []string `json:"cartDateAdded"`
}

// CartUploadResponse struct
type CartSyncResponse struct {
	Status string `json:"status"`
}

func CartSync(cartSyncRequest *CartSyncRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	for i := 0; i < len(cartSyncRequest.GameID); i++ {

		cart := new(model.Cart)
		cart.UserID = cartSyncRequest.UserID
		cart.GameID = cartSyncRequest.GameID[i]
		cart.CartDateAdded = cartSyncRequest.CartDateAdded[i]

		// check inventory
		var count1 int64
		var count2 int64
		if err := tx.Debug().Where("user_id = ? AND game_id = ?", cart.UserID, cart.GameID).Find(&model.Inventory{}).Count(&count1).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Debug().Where("user_id = ? AND game_id = ?", cart.UserID, cart.GameID).Find(&model.Cart{}).Count(&count2).Error; err != nil {
			tx.Rollback()
			return err
		}

		if count1 == 0 && count2 == 0 {
			// Upload games to a cart
			if err := tx.Create(&cart).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

	}

	tx.Commit()
	return tx.Error
}

func CartSyncOutput(w http.ResponseWriter, cartSyncResponse *CartSyncResponse) {
	jsonbyte, err := json.Marshal(cartSyncResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
