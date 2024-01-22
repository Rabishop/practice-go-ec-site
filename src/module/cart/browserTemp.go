package cart

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// CartBrowerTempRequest struct
type CartBrowserTempRequest struct {
	GameID []int64 `json:"gameID"`
}

// CartBrowerTempResponse struct
type CartBrowserTempResponse struct {
	Status   string       `json:"status"`
	GameList []model.Game `json:"gameList"`
}

func CartBrowserTemp(cartBrowserTempRequest *CartBrowserTempRequest, cartBrowserTempResponse *CartBrowserTempResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	gameItem := new(model.Game)

	// get gameList by gameID
	for i := 0; i < len(cartBrowserTempRequest.GameID); i++ {
		if err := tx.Where("game_id = ?", cartBrowserTempRequest.GameID[i]).Take(&gameItem).Error; err != nil {
			tx.Rollback()
			return err
		}
		// fmt.Println(gameItem)
		cartBrowserTempResponse.GameList = append(cartBrowserTempResponse.GameList, *gameItem)
	}

	// fmt.Println(cartBrowerResponse)

	tx.Commit()
	return tx.Error
}

func CartBrowserTempOutput(w http.ResponseWriter, cartBrowserTempResponse *CartBrowserTempResponse) {
	jsonbyte, err := json.Marshal(cartBrowserTempResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
