package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		state, _ := c.Cookie("oauth_state")
		if state != c.Query("state") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OAuth state"})
			return
		}
		c.SetCookie("oauth_state", "", -1, "/", "", true, true)

		code := c.Query("code")
		token, err := oauth.GoogleOAuthConfig.Exchange(c, code)
		if err != nil {
			logging.LogError(err, map[string]interface{}{
				"event": "google_oauth_exchange_failed",
			})
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange token"})
			return
		}

		client := oauth.GoogleOAuthConfig.Client(c, token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			logging.LogError(err, map[string]interface{}{
				"event": "google_userinfo_failed",
			})
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
			return
		}
		defer resp.Body.Close()

		userBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logging.LogError(err, map[string]interface{}{
				"event": "google_userinfo_read_failed",
			})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user info"})
			return
		}

		var googleUser GoogleUserInfo
		if err := json.Unmarshal(userBytes, &googleUser); err != nil {
			logging.LogError(err, map[string]interface{}{
				"event": "google_userinfo_parse_failed",
			})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}

			user = User{
				Usr:  googleUser.Email,
				Pwd:  string(hashedPwd),
				Role: RoleUsr,
			}

			if err := db.Create(&user).Error; err != nil {
				logging.LogError(err, map[string]interface{}{
					"event": "user_creation_failed",
					"email": googleUser.Email,
				})
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}
		} else if result.Error != nil {
			logging.LogError(result.Error, map[string]interface{}{
				"event": "user_lookup_failed",
				"email": googleUser.Email,
			})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process user"})
			return
		}

		jwtToken, err := genUsrToken(c, &user)
		if err != nil {
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

		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"id":       user.ID,
				"username": user.Usr,
				"role":     user.Role,
			},
		})
	}
}
