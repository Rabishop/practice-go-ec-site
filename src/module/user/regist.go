package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/v2/model"
	"golang.org/x/crypto/bcrypt"
)

// UserRegistResponse struct
type UserRegistResponse struct {
	Status string `json:"status"`
}

// UserRegistRequest struct
type UserRegistRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	UserName     string `json:"userName"`
}

// 生成密码的哈希值
func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Regist(userRegistRequest *UserRegistRequest) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	user.UserName = userRegistRequest.UserName
	user.UserAccount = userRegistRequest.UserAccount
	if passwordHash, err := hashPassword(userRegistRequest.UserPassword); err != nil {
		log.Println(err)
	} else {
		user.UserPassword = string(passwordHash)
	}

	// Upload a new user
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return tx.Error
}

func RegistOutput(w http.ResponseWriter, userRegistResponse *UserRegistResponse) {
	jsonbyte, err := json.Marshal(userRegistResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
