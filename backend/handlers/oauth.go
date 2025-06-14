package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"vhu_bmwa/logging"
	"vhu_bmwa/oauth"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

func GoogleLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		state := fmt.Sprintf("%d", getTymNow().UnixNano())
		c.SetCookie("oauth_state", state, 3600, "/", "", true, true)

		url := oauth.GoogleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
		c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func GoogleCallbackHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		frontendURL := os.Getenv("FRONTEND_URL")
		if frontendURL == "" {
			frontendURL = "http://localhost:5173"
		}

		state, _ := c.Cookie("oauth_state")
		if state != c.Query("state") {
			c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=invalid_state")
			return
		}
		c.SetCookie("oauth_state", "", -1, "/", "", true, true)

		code := c.Query("code")
		token, err := oauth.GoogleOAuthConfig.Exchange(c, code)
		if err != nil {
			logging.LogError(err, map[string]interface{}{
				"event": "google_oauth_exchange_failed",
			})
			c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=token_exchange_failed")
			return
		}

		client := oauth.GoogleOAuthConfig.Client(c, token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			logging.LogError(err, map[string]interface{}{
				"event": "google_userinfo_failed",
			})
			c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=userinfo_failed")
			return
		}
		defer resp.Body.Close()

		userBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logging.LogError(err, map[string]interface{}{
				"event": "google_userinfo_read_failed",
			})
			c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=userinfo_read_failed")
			return
		}

		var googleUser GoogleUserInfo
		if err := json.Unmarshal(userBytes, &googleUser); err != nil {
			logging.LogError(err, map[string]interface{}{
				"event": "google_userinfo_parse_failed",
			})
			c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=userinfo_parse_failed")
			return
		}

		var user User
		result := db.Where("usr = ?", googleUser.Email).First(&user)
		if result.Error == gorm.ErrRecordNotFound {
			hashedPwd, err := hashPwd(googleUser.ID)
			if err != nil {
				logging.LogError(err, map[string]interface{}{
					"event": "password_hash_failed",
					"email": googleUser.Email,
				})
				c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=user_creation_failed")
				return
			}

			user = User{
				Username: googleUser.Email,
				Password: string(hashedPwd),
				Role:     RoleUsr,
			}

			if err := db.Create(&user).Error; err != nil {
				logging.LogError(err, map[string]interface{}{
					"event": "user_creation_failed",
					"email": googleUser.Email,
				})
				c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=user_creation_failed")
				return
			}
		} else if result.Error != nil {
			logging.LogError(result.Error, map[string]interface{}{
				"event": "user_lookup_failed",
				"email": googleUser.Email,
			})
			c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=user_lookup_failed")
			return
		}

		jwtToken, err := genUsrToken(c, &user)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/login?error=token_generation_failed")
			return
		}

		setAuthCookie(c, jwtToken)

		logging.LogSecurity(logging.SecurityLogEntry{
			EventType:   "GOOGLE_LOGIN_SUCCESS",
			UserID:      user.ID,
			IP:          c.ClientIP(),
			Description: "User logged in via Google OAuth",
			Metadata: map[string]interface{}{
				"email": googleUser.Email,
			},
		})

		// Redirect to frontend home page after successful login
		c.Redirect(http.StatusTemporaryRedirect, frontendURL)
	}
}
