package game

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
)

// GameIndexResquest struct
type GameIndexResquest struct {
}

// GameIndexResponse struct
type GameIndexResponse struct {
	Status    string           `json:"status"`
	GameIndex []GameList       `json:"gameIndex"`
	GameType  []model.Tag_list `json:"gameType"`
}

type GameList struct {
	GameItem []model.Game `json:"gameItem"`
}

func GameIndex(gameIndexResponse *GameIndexResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	// get games on the homepage
	gameList := make([]model.Tag, 100, 200)

	if err := tx.Raw("(SELECT * FROM tag AS a WHERE (SELECT COUNT(*) FROM tag AS b WHERE b.tag_name = a.tag_name AND b.game_id >= a.game_id) <= 4 ORDER BY a.game_id ASC) ORDER BY tag_id").Scan(&gameList).Error; err != nil {
		tx.Rollback()
		return err
	}

	tagNow := ""
	tagCount := 0
	var tempList [100]GameList
	for i := 0; i < len(gameList); i++ {

		if tagNow != gameList[i].TagName || i == (len(gameList)-1) {
			tagNow = gameList[i].TagName

			if i != 0 {
				gameIndexResponse.GameIndex = append(gameIndexResponse.GameIndex, tempList[tagCount])
			}
			tagCount++
		}

		item := new(model.Game)
		if err := tx.Where("game_id = ?", gameList[i].GameID).Find(item).Error; err != nil {
			tx.Rollback()
			return err
		}

		tempList[tagCount].GameItem = append(tempList[tagCount].GameItem, *item)

		if i == (len(gameList) - 1) {
			gameIndexResponse.GameIndex = append(gameIndexResponse.GameIndex, tempList[tagCount])
		}
	}

	if err := tx.Order("tag_id asc").Find(&gameIndexResponse.GameType).Error; err != nil {
		tx.Rollback()
		return err
	}

	// fmt.Println(gameIndexResponse.GameType)

	tx.Commit()
	return tx.Error
}

func GameIndexOutput(w http.ResponseWriter, gameIndexResponse *GameIndexResponse) {
	jsonbyte, err := json.Marshal(gameIndexResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
