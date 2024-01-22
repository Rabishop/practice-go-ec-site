package user

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/m/v2/model"
	"golang.org/x/crypto/bcrypt"
)

// UserLoginResponse struct
type UserLoginResponse struct {
	UserSessionID string `json:"userSessionID"`
	Status        string `json:"status"`
}

// UserRegistRequest struct
type UserLoginRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}

func checkPasswordHash(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func Login(userLoginRequest *UserLoginRequest, userLoginResponse *UserLoginResponse) error {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	var user model.User
	user.UserAccount = userLoginRequest.UserAccount

	// user login
	if err := tx.Where("user_account = ?", user.UserAccount).Take(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// check user's password with hash
	if flag := checkPasswordHash([]byte(userLoginRequest.UserPassword), []byte(user.UserPassword)); !flag {
		return errors.New("password error")
	}

	// create and store user's sessionID
	var session model.Session
	session.SessionID = CreateSessionID()
	session.UserID = user.UserID
	session.SessionLastTime = time.Now().String()[0:19]

	// Upload a new user's session
	if err := tx.Where("user_id = ?", session.UserID).Save(&session).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	userLoginResponse.UserSessionID = session.SessionID

	// // check the session
	// if flag, err := checkSessionID(session.SessionID); err != nil {
	// 	return err
	// }

	return tx.Error
}

func CreateSessionID() string {
	// Generate a random session ID
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	// Hash the session ID using SHA-256
	h := sha256.Sum256(b)

	// Encode the hashed session ID in base64
	sessionID := base64.StdEncoding.EncodeToString(h[:])

	// fmt.Println(sessionID)

	return sessionID
}

func CheckSessionID(sessionID string) (int64, error) {

	// connect database
	DB := model.MysqlConn()

	// start transcation
	tx := DB.Begin()

	// 检查session ID是否为空
	if sessionID == "" {
		return 0, nil
	}

	// 查询数据库，检查session ID是否存在
	var count int64
	var session model.Session
	if err := tx.Where("session_id = ?", sessionID).Find(&session).Count(&count).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if count == 0 {
		return 0, errors.New("no such sessionID")
	} else {
		return session.UserID, nil
	}
}

func LoginOutput(w http.ResponseWriter, userLoginResponse *UserLoginResponse) {
	jsonbyte, err := json.Marshal(userLoginResponse)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonbyte))
}
