package cart

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/model"
)

// CartUploadRequest struct
type CartUploadRequest struct {
	UserID        int64  `json:"userID"`
	GameID        int64  `json:"gameID"`
	CartDateAdded string `json:"cartDateAdded"`
}

// CartUploadResponse struct
type CartUploadResponse struct {
	Status string `json:"status"`
}

func CartUpload(cartUploadRequest *CartUploadRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	cart := new(model.Cart)
	cart.UserID = cartUploadRequest.UserID
	cart.GameID = cartUploadRequest.GameID
	cart.CartDateAdded = time.Now().String()[0:20]

	// check inventory
	var count int64
	if err := tx.Where("user_id = ? AND game_id = ?", cart.UserID, cartUploadRequest.GameID).Find(&model.Inventory{}).Count(&count).Error; err != nil {
		tx.Rollback()
		return err
	}

	if count > 0 {
		return errors.New("already in the inventory")
	}

	// Upload games to a cart
	if err := tx.Create(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func CartUploadOutput(w http.ResponseWriter, cartUploadResponse *CartUploadResponse) {
	jsonbyte, err := json.Marshal(cartUploadResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
