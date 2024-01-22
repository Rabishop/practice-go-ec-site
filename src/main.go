package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/m/v2/module/cart"
	"example.com/m/v2/module/game"
	"example.com/m/v2/module/user"
)

// UserRegist
func UserRegistHandler(w http.ResponseWriter, r *http.Request) {

	var userRegistRequest user.UserRegistRequest
	var userRegistResponse user.UserRegistResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userRegistResponse.Status = "Wrong Method"
		user.RegistOutput(w, &userRegistResponse)
		return
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &userRegistRequest)

	// Call Function
	err = user.Regist(&userRegistRequest)
	if err != nil {
		log.Println(err)
		userRegistResponse.Status = "SQL Access Error"
		user.RegistOutput(w, &userRegistResponse)
		return
	}

	// Return JSON
	userRegistResponse.Status = "Accepted"
	user.RegistOutput(w, &userRegistResponse)
}

// UserLogin
func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	var userLoginRequest user.UserLoginRequest
	var userLoginResponse user.UserLoginResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userLoginResponse.Status = "Wrong Method"
		user.LoginOutput(w, &userLoginResponse)
	}
	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal([]byte(body), &userLoginRequest)

	// Call Function
	err = user.Login(&userLoginRequest, &userLoginResponse)
	if err != nil {
		log.Println(err)
		userLoginResponse.Status = "SQL Access Error"
		user.LoginOutput(w, &userLoginResponse)
		return
	}

	// Set Cookie
	cookie := http.Cookie{Name: "sessionID", Value: userLoginResponse.UserSessionID, Path: "/", MaxAge: 86400}
	http.SetCookie(w, &cookie)

	// Return JSON
	userLoginResponse.Status = "Accepted"
	user.LoginOutput(w, &userLoginResponse)
}

// UserLogout
func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {

	var userLogoutRequest user.UserLogoutRequest
	var userLogoutResponse user.UserLoginResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userLogoutResponse.Status = "Wrong Method"
		user.LogoutOutput(w, &userLogoutResponse)
	}

	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		userLogoutResponse.Status = "Cookie Error"
		user.LogoutOutput(w, &userLogoutResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			return
		} else {
			userLogoutRequest.UserID = userID
		}
	}

	// Call Function
	err := user.Logout(&userLogoutRequest)
	if err != nil {
		log.Println(err)
		userLogoutResponse.Status = "SQL Access Error"
		user.LogoutOutput(w, &userLogoutResponse)
		return
	}

	// Return JSON
	userLogoutResponse.Status = "Accepted"
	user.LogoutOutput(w, &userLogoutResponse)
}

// View user's Profile
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {

	var userProfileRequest user.UserProfileRequest
	var userProflieResponse user.UserProfileResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userProflieResponse.Status = "Wrong Method"
		user.ProfileOutput(w, &userProflieResponse)
	}
	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		userProflieResponse.Status = "Cookie Error"
		user.ProfileOutput(w, &userProflieResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			userProflieResponse.Status = "Cookie Error"
			user.ProfileOutput(w, &userProflieResponse)
			return
		} else {
			userProfileRequest.UserID = userID
		}
	}

	// Call Function
	err := user.Profile(&userProfileRequest, &userProflieResponse)
	if err != nil {
		log.Println(err)
		userProflieResponse.Status = "SQL Access Error"
		user.ProfileOutput(w, &userProflieResponse)
		return
	}

	// Return JSON
	userProflieResponse.Status = "Accepted"
	user.ProfileOutput(w, &userProflieResponse)
}

// View user's inventory
func UserInventoryHandler(w http.ResponseWriter, r *http.Request) {

	var userInventoryRequest user.UserInventoryRequest
	var userInventoryResponse user.UserInventoryResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userInventoryResponse.Status = "Wrong Method"
		user.InventoryOutput(w, &userInventoryResponse)
	}
	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		userInventoryResponse.Status = "Cookie Error"
		user.InventoryOutput(w, &userInventoryResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			return
		} else {
			userInventoryRequest.UserID = userID
		}
	}

	// Call Function
	err := user.Inventory(&userInventoryRequest, &userInventoryResponse)
	if err != nil {
		log.Println(err)
		userInventoryResponse.Status = "SQL Access Error"
		user.InventoryOutput(w, &userInventoryResponse)
		return
	}

	// Return JSON
	userInventoryResponse.Status = "Accepted"
	user.InventoryOutput(w, &userInventoryResponse)
}

// Upload user's portrait
func UserUploadPortraitHandler(w http.ResponseWriter, r *http.Request) {

	var userUploadPortraitRequest user.UserUploadPortraitRequest
	var userUploadPortraitResponse user.UserUploadPortraitResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		userUploadPortraitResponse.Status = "Wrong Method"
		user.UploadPortraitOutput(w, &userUploadPortraitResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &userUploadPortraitRequest)

	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		userUploadPortraitResponse.Status = "Cookie Error"
		user.UploadPortraitOutput(w, &userUploadPortraitResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			return
		} else {
			userUploadPortraitRequest.UserID = userID
		}
	}

	// Call Function
	err = user.UploadPortrait(&userUploadPortraitRequest)
	if err != nil {
		log.Println(err)
		userUploadPortraitResponse.Status = "SQL Access Error"
		user.UploadPortraitOutput(w, &userUploadPortraitResponse)
		return
	}

	// Return JSON
	userUploadPortraitResponse.Status = "Accepted"
	user.UploadPortraitOutput(w, &userUploadPortraitResponse)
}

// Upload games
func GameUploadHandler(w http.ResponseWriter, r *http.Request) {

	var gameUploadRequest game.GameUploadRequest
	var gameUploadResponse game.GameUploadResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		gameUploadResponse.Status = "Wrong Method"
		game.GameUploadOutput(w, &gameUploadResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &gameUploadRequest)

	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		gameUploadResponse.Status = "Cookie Error"
		game.GameUploadOutput(w, &gameUploadResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			gameUploadResponse.Status = "Cookie Error"
			game.GameUploadOutput(w, &gameUploadResponse)
			return
		} else {
			gameUploadRequest.UserID = userID
		}
	}

	// Call Function
	err = game.GameUpload(&gameUploadRequest)
	if err != nil {
		log.Println(err)
		gameUploadResponse.Status = "SQL Access Error"
		game.GameUploadOutput(w, &gameUploadResponse)
		return
	}

	// Return JSON
	gameUploadResponse.Status = "Accepted"
	game.GameUploadOutput(w, &gameUploadResponse)
}

// View games on the homepage
func GameIndexHandler(w http.ResponseWriter, r *http.Request) {

	var gameIndexResponse game.GameIndexResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		gameIndexResponse.Status = "Wrong Method"
		game.GameIndexOutput(w, &gameIndexResponse)
	}

	// Call Function
	err := game.GameIndex(&gameIndexResponse)
	if err != nil {
		log.Println(err)
		gameIndexResponse.Status = "SQL Access Error"
		game.GameIndexOutput(w, &gameIndexResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	gameIndexResponse.Status = "Accepted"
	game.GameIndexOutput(w, &gameIndexResponse)
}

// View games by types
func GameBrowserHandler(w http.ResponseWriter, r *http.Request) {

	var gameBrowserResqust game.GameBrowserResqust
	var gameBrowserResponse game.GameBrowserResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		gameBrowserResponse.Status = "Wrong Method"
		game.GameBrowserOutput(w, &gameBrowserResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &gameBrowserResqust)

	// Call Function
	err = game.GameBrowser(&gameBrowserResqust, &gameBrowserResponse)
	if err != nil {
		log.Println(err)
		gameBrowserResponse.Status = "SQL Access Error"
		game.GameBrowserOutput(w, &gameBrowserResponse)
		return
	}

	// Return JSON
	gameBrowserResponse.Status = "Accepted"
	game.GameBrowserOutput(w, &gameBrowserResponse)
}

// View game's details
func GameDetailsHandler(w http.ResponseWriter, r *http.Request) {

	var gameDetailsRequest game.GameDetailsRequest
	var gameDetailsResponse game.GameDetailsResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		gameDetailsResponse.Status = "Wrong Method"
		game.GameDetailsOutput(w, &gameDetailsResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &gameDetailsRequest)

	// check cookie (can access game info without session)
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
		} else {
			gameDetailsRequest.UserID = userID
		}
	}

	// Call Function
	err = game.GameDetails(&gameDetailsRequest, &gameDetailsResponse)
	if err != nil {
		log.Println(err)
		gameDetailsResponse.Status = "SQL Access Error"
		game.GameDetailsOutput(w, &gameDetailsResponse)
		return
	}

	// Return JSON
	gameDetailsResponse.Status = "Accepted"

	game.GameDetailsOutput(w, &gameDetailsResponse)
}

// Add games to a cart
func CartUploadHandler(w http.ResponseWriter, r *http.Request) {

	var cartUploadRequest cart.CartUploadRequest
	var cartUploadResponse cart.CartUploadResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartUploadResponse.Status = "Wrong Method"
		cart.CartUploadOutput(w, &cartUploadResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartUploadRequest)

	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		cartUploadResponse.Status = "Cookie Error"
		cart.CartUploadOutput(w, &cartUploadResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			cartUploadResponse.Status = "Cookie Error"
			cart.CartUploadOutput(w, &cartUploadResponse)
			return
		} else {
			cartUploadRequest.UserID = userID
		}
	}

	// Call Function
	err = cart.CartUpload(&cartUploadRequest)
	if err != nil {
		log.Println(err)
		cartUploadResponse.Status = "SQL Access Error"
		cart.CartUploadOutput(w, &cartUploadResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartUploadResponse.Status = "Accepted"
	cart.CartUploadOutput(w, &cartUploadResponse)
}

// View the cart
func CartBrowserHandler(w http.ResponseWriter, r *http.Request) {

	var cartBrowserResqust cart.CartBrowserResqust
	var cartBrowserResponse cart.CartBrowserResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartBrowserResponse.Status = "Wrong Method"
		cart.CartBrowserOutput(w, &cartBrowserResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartBrowserResqust)

	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		cartBrowserResponse.Status = "Cookie Error"
		cart.CartBrowserOutput(w, &cartBrowserResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			cartBrowserResponse.Status = "Cookie Error"
			cart.CartBrowserOutput(w, &cartBrowserResponse)
			return
		} else {
			cartBrowserResqust.UserID = userID
		}
	}

	// Call Function
	err = cart.CartBrowser(&cartBrowserResqust, &cartBrowserResponse)
	if err != nil {
		log.Println(err)
		cartBrowserResponse.Status = "SQL Access Error"
		cart.CartBrowserOutput(w, &cartBrowserResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartBrowserResponse.Status = "Accepted"
	cart.CartBrowserOutput(w, &cartBrowserResponse)
}

// View the cart
func CartBrowserTempHandler(w http.ResponseWriter, r *http.Request) {

	var cartBrowserTempResqust cart.CartBrowserTempRequest
	var cartBrowserTempResponse cart.CartBrowserTempResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartBrowserTempResponse.Status = "Wrong Method"
		cart.CartBrowserTempOutput(w, &cartBrowserTempResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartBrowserTempResqust)

	// fmt.Println(cartBrowserTempResqust)

	// Call Function
	err = cart.CartBrowserTemp(&cartBrowserTempResqust, &cartBrowserTempResponse)
	if err != nil {
		log.Println(err)
		cartBrowserTempResponse.Status = "SQL Access Error"
		cart.CartBrowserTempOutput(w, &cartBrowserTempResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartBrowserTempResponse.Status = "Accepted"
	cart.CartBrowserTempOutput(w, &cartBrowserTempResponse)
}

// Remove games from a cart
func CartRemoveHandler(w http.ResponseWriter, r *http.Request) {

	var cartRemoveRequest cart.CartRemoveRequest
	var cartRemoveResponse cart.CartRemoveResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartRemoveResponse.Status = "Wrong Method"
		cart.CartRemoveOutput(w, &cartRemoveResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartRemoveRequest)

	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		cartRemoveResponse.Status = "Cookie Error"
		cart.CartRemoveOutput(w, &cartRemoveResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			cartRemoveResponse.Status = "Cookie Error"
			cart.CartRemoveOutput(w, &cartRemoveResponse)
			return
		} else {
			cartRemoveRequest.UserID = userID
		}
	}

	// Call Function
	err = cart.CartRemove(&cartRemoveRequest)
	if err != nil {
		log.Println(err)
		cartRemoveResponse.Status = "SQL Access Error"
		cart.CartRemoveOutput(w, &cartRemoveResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartRemoveResponse.Status = "Accepted"
	cart.CartRemoveOutput(w, &cartRemoveResponse)
}

// Cart synchronization
func CartSyncHandler(w http.ResponseWriter, r *http.Request) {

	var cartSyncRequest cart.CartSyncRequest
	var cartSyncResponse cart.CartSyncResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartSyncResponse.Status = "Wrong Method"
		cart.CartSyncOutput(w, &cartSyncResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartSyncRequest)

	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		cartSyncResponse.Status = "Cookie Error"
		cart.CartSyncOutput(w, &cartSyncResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			cartSyncResponse.Status = "Cookie Error"
			cart.CartSyncOutput(w, &cartSyncResponse)
			return
		} else {
			cartSyncRequest.UserID = userID
		}
	}

	// fmt.Println(cartSyncRequest.GameID)
	// fmt.Println(cartSyncRequest.CartDateAdded)

	// Call Function
	if len(cartSyncRequest.GameID) > 0 {
		err = cart.CartSync(&cartSyncRequest)
		if err != nil {
			log.Println(err)
			cartSyncResponse.Status = "SQL Access Error"
			cart.CartSyncOutput(w, &cartSyncResponse)
			return
		}
		cartSyncResponse.Status = "CartSync"
		cart.CartSyncOutput(w, &cartSyncResponse)
	} else {
		// Return JSON
		cartSyncResponse.Status = "Accepted"
		cart.CartSyncOutput(w, &cartSyncResponse)
	}

}

// Add games to a inventory and Remove games from the cart
func CartCheckHandler(w http.ResponseWriter, r *http.Request) {

	var cartCheckRequest cart.CartCheckRequest
	var cartCheckResponse cart.CartCheckResponse

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	// Check Method
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		cartCheckResponse.Status = "Wrong Method"
		cart.CartCheckOutput(w, &cartCheckResponse)
	}

	// Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal([]byte(body), &cartCheckRequest)

	// check cookie
	if cookie, err := r.Cookie("sessionID"); err != nil {
		log.Println(err)
		cartCheckResponse.Status = "Cookie Error"
		cart.CartCheckOutput(w, &cartCheckResponse)
		return
	} else {
		// check the session
		if userID, err := user.CheckSessionID(cookie.Value); err != nil {
			log.Println(err)
			cartCheckResponse.Status = "Cookie Error"
			cart.CartCheckOutput(w, &cartCheckResponse)
			return
		} else {
			cartCheckRequest.UserID = userID
		}
	}

	// fmt.Println(cartRemoveRequest)

	// Call Function
	err = cart.CartCheck(&cartCheckRequest)
	if err != nil {
		log.Println(err)
		cartCheckResponse.Status = "SQL Access Error"
		cart.CartCheckOutput(w, &cartCheckResponse)
		return
	}

	// fmt.Println(gameIndexResponse)

	// Return JSON
	cartCheckResponse.Status = "Accepted"
	cart.CartCheckOutput(w, &cartCheckResponse)
}

func main() {

	// Functions Handle

	http.HandleFunc("/user/regist", UserRegistHandler)

	http.HandleFunc("/user/login", UserLoginHandler)

	http.HandleFunc("/user/logout", UserLogoutHandler)

	http.HandleFunc("/user/profile", UserProfileHandler)

	http.HandleFunc("/user/inventory", UserInventoryHandler)

	http.HandleFunc("/user/uploadPortrait", UserUploadPortraitHandler)

	http.HandleFunc("/game/upload", GameUploadHandler)

	http.HandleFunc("/game/index", GameIndexHandler)

	http.HandleFunc("/game/browser", GameBrowserHandler)

	http.HandleFunc("/game/details", GameDetailsHandler)

	http.HandleFunc("/cart/upload", CartUploadHandler)

	http.HandleFunc("/cart/browser", CartBrowserHandler)

	http.HandleFunc("/cart/browserTemp", CartBrowserTempHandler)

	http.HandleFunc("/cart/remove", CartRemoveHandler)

	http.HandleFunc("/cart/sync", CartSyncHandler)

	http.HandleFunc("/cart/check", CartCheckHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../pages/assets"))))

	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("../pages/vendor"))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../pages/"+r.URL.Path+".html")
	}))

	// Build the Server

	http.ListenAndServe(":8080", nil)
}
