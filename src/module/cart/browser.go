package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// CartBrowerResponse struct
type CartBrowserResponse struct {
	Status   string       `json:"status"`
	CartList []model.Cart `json:"cartList"`
	GameList []model.Game `json:"gameList"`
}

// CartBrowerRequest struct
type CartBrowserResqust struct {
	UserID int64 `json:"userID"`
}

func CartBrowser(cartBrowerResqust *CartBrowserResqust, cartBrowerResponse *CartBrowserResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	gameItem := new(model.Game)

	//　get gameID in the cart
	if err := tx.Where("user_id = ?", cartBrowerResqust.UserID).Find(&cartBrowerResponse.CartList).Error; err != nil {
		tx.Rollback()
		return err
	}

	// get gameList by gameID
	for i := 0; i < len(cartBrowerResponse.CartList); i++ {
		if err := tx.Where("game_id = ?", cartBrowerResponse.CartList[i].GameID).Take(&gameItem).Error; err != nil {
			tx.Rollback()
			return err
		}
		// fmt.Println(gameItem)
		cartBrowerResponse.GameList = append(cartBrowerResponse.GameList, *gameItem)
	}

	// fmt.Println(cartBrowerResponse)

	tx.Commit()
	return tx.Error
}

func CartBrowserOutput(w http.ResponseWriter, cartBrowerResponse *CartBrowserResponse) {
	jsonbyte, err := json.Marshal(cartBrowerResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
