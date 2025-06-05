package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode"
	"vhu_bmwa/logging"
	"vhu_bmwa/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	MinPwdLen = 8
	MaxPwdLen = 72
	MinUsrLen = 3
	MaxUsrLen = 32
)

const (
	RoleUsr = "user"
	RoleAdm = "admin"
	RoleMod = "moderator"
)

type User struct {
	gorm.Model
	Usr  string `json:"username" gorm:"unique;not null"`
	Pwd  string `json:"-" gorm:"not null"`
	Role string `json:"role" gorm:"default:user"`
}

type LoginReq struct {
	Usr string `json:"username" binding:"required"`
	Pwd string `json:"password" binding:"required"`
}

type RegReq struct {
	Usr string `json:"username" binding:"required"`
	Pwd string `json:"password" binding:"required"`
}

type ChgPwdReq struct {
	CurPwd string `json:"currentPassword" binding:"required"`
	NewPwd string `json:"newPassword" binding:"required"`
}

func chekPwdLen(p string) bool {
	return len(p) >= MinPwdLen && len(p) <= MaxPwdLen
}

func chekUsrLen(u string) bool {
	return len(u) >= MinUsrLen && len(u) <= MaxUsrLen
}

func chekPwdChars(p string) (bool, bool, bool, bool) {
	var hasUp, hasLow, hasNum, hasSym bool = false, false, false, false
	for _, c := range p {
		switch {
		case unicode.IsUpper(c):
			hasUp = true
		case unicode.IsLower(c):
			hasLow = true
		case unicode.IsNumber(c):
			hasNum = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSym = true
		default:
			continue
		}
	}
	return hasUp, hasLow, hasNum, hasSym
}

func chekUsrChars(u string) bool {
	validUsr := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	result := validUsr.MatchString(u)
	return result
}

func getTymNow() time.Time {
	currentTym := time.Now()
	return currentTym
}

func genErMsg(msg string) string {
	erMsg := fmt.Sprintf("error: %s", msg)
	return erMsg
}

func ValidatePwd(p string) error {
	if !chekPwdLen(p) {
		return fmt.Errorf(genErMsg(fmt.Sprintf("password length must be between %d and %d", MinPwdLen, MaxPwdLen)))
	}
	hasUp, hasLow, hasNum, hasSym := chekPwdChars(p)

	var missingChars []string
	if !hasUp {
		missingChars = append(missingChars, "uppercase")
	}
	if !hasLow {
		missingChars = append(missingChars, "lowercase")
	}
	if !hasNum {
		missingChars = append(missingChars, "number")
	}
	if !hasSym {
		missingChars = append(missingChars, "special")
	}

	if len(missingChars) > 0 {
		return fmt.Errorf(genErMsg(fmt.Sprintf("password must contain: %s", strings.Join(missingChars, ", "))))
	}

	return nil
}

func ValidateUsr(u string) error {
	if !chekUsrLen(u) {
		return fmt.Errorf(genErMsg(fmt.Sprintf("username length must be between %d and %d", MinUsrLen, MaxUsrLen)))
	}

	if !chekUsrChars(u) {
		return fmt.Errorf(genErMsg("username can only contain letters, numbers, and underscores"))
	}

	return nil
}

func setAuthCookie(c *gin.Context, token string) {
	c.SetCookie(
		"auth_token",
		token,
		24*60*60, //(24 hours in seconds)
		"/",
		"",
		true,
		true,
	)
}

func RegHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegReq
		if err := parseRegReq(c, &req); err != nil {
			handleRegErr(c, "PARSE_FAILED", err, nil)
			return
		}

		if err := validateRegUsr(c, db, req.Usr); err != nil {
			return
		}

		if err := validateRegPwd(c, req.Pwd); err != nil {
			return
		}

		usr, err := createUsr(db, req.Usr, req.Pwd)
		if err != nil {
			handleRegErr(c, "CREATE_FAILED", err, map[string]interface{}{
				"username": req.Usr,
			})
			return
		}

		token, err := genUsrToken(c, usr)
		if err != nil {
			return
		}

		setAuthCookie(c, token)
		logRegSuccess(c, usr)

		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"id":       usr.ID,
				"username": usr.Usr,
				"role":     usr.Role,
			},
		})
	}
}

func parseRegReq(c *gin.Context, req *RegReq) error {
	return c.ShouldBindJSON(req)
}

func validateRegUsr(c *gin.Context, db *gorm.DB, username string) error {
	if err := ValidateUsr(username); err != nil {
		handleRegErr(c, "USERNAME_INVALID", err, map[string]interface{}{
			"username": username,
		})
		return err
	}

	var existingUsr User = User{}
	if result := db.Where("usr = ?", username).First(&existingUsr); result.Error == nil {
		err := fmt.Errorf("username already exists")
		handleRegErr(c, "USERNAME_DUPLICATE", err, map[string]interface{}{
			"username": username,
		})
		return err
	}

	return nil
}

func validateRegPwd(c *gin.Context, password string) error {
	if err := ValidatePwd(password); err != nil {
		handleRegErr(c, "PASSWORD_INVALID", err, map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func createUsr(db *gorm.DB, username, password string) (*User, error) {
	hashedPwd, err := hashPwd(password)
	if err != nil {
		return nil, err
	}

	usr := &User{
		Usr:  username,
		Pwd:  string(hashedPwd),
		Role: RoleUsr,
	}

	if err := db.Create(usr).Error; err != nil {
		return nil, err
	}

	return usr, nil
}

func hashPwd(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func genUsrToken(c *gin.Context, usr *User) (string, error) {
	token, err := middleware.GenerateToken(usr.ID, usr.Role)
	if err != nil {
		handleRegErr(c, "TOKEN_FAILED", err, map[string]interface{}{
			"user_id": usr.ID,
		})
		return "", err
	}
	return token, nil
}

func handleRegErr(c *gin.Context, eventType string, err error, metadata map[string]interface{}) {
	logging.LogSecurity(logging.SecurityLogEntry{
		EventType:   fmt.Sprintf("REGISTER_%s", eventType),
		IP:          c.ClientIP(),
		Description: err.Error(),
		Metadata:    metadata,
	})

	c.JSON(getErrStatusCode(eventType), gin.H{
		"error": err.Error(),
	})
}

func getErrStatusCode(eventType string) int {
	switch eventType {
	case "PARSE_FAILED", "USERNAME_INVALID", "PASSWORD_INVALID":
		return http.StatusBadRequest
	case "USERNAME_DUPLICATE":
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func logRegSuccess(c *gin.Context, usr *User) {
	logging.LogSecurity(logging.SecurityLogEntry{
		EventType:   "REGISTER_SUCCESS",
		UserID:      usr.ID,
		IP:          c.ClientIP(),
		Description: "User registered successfully",
		Metadata: map[string]interface{}{
			"username": usr.Usr,
			"role":     usr.Role,
		},
	})
}

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginReq
		if err := parseLoginReq(c, &req); err != nil {
			handleLoginErr(c, "PARSE_FAILED", err, nil)
			return
		}

		usr, err := getUsr(db, req.Usr)
		if err != nil {
			handleLoginErr(c, "USER_NOT_FOUND", err, map[string]interface{}{
				"username": req.Usr,
			})
			return
		}

		if err := chekUsrPwd(usr, req.Pwd); err != nil {
			handleLoginErr(c, "PASSWORD_INVALID", err, map[string]interface{}{
				"username": req.Usr,
			})
			return
		}

		token, err := genUsrToken(c, usr)
		if err != nil {
			return
		}

		setAuthCookie(c, token)
		logLoginSuccess(c, usr)

		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"id":       usr.ID,
				"username": usr.Usr,
				"role":     usr.Role,
			},
		})
	}
}

func parseLoginReq(c *gin.Context, req *LoginReq) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}

func getUsr(db *gorm.DB, username string) (*User, error) {
	var usr User
	if err := db.Where("usr = ?", username).First(&usr).Error; err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}
	return &usr, nil
}

func chekUsrPwd(usr *User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(usr.Pwd), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}

func handleLoginErr(c *gin.Context, eventType string, err error, metadata map[string]interface{}) {
	logging.LogSecurity(logging.SecurityLogEntry{
		EventType:   fmt.Sprintf("LOGIN_%s", eventType),
		IP:          c.ClientIP(),
		Description: err.Error(),
		Metadata:    metadata,
	})

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid username or password",
	})
}

func logLoginSuccess(c *gin.Context, usr *User) {
	logging.LogSecurity(logging.SecurityLogEntry{
		EventType:   "LOGIN_SUCCESS",
		UserID:      usr.ID,
		IP:          c.ClientIP(),
		Description: "User logged in successfully",
		Metadata: map[string]interface{}{
			"username": usr.Usr,
			"role":     usr.Role,
		},
	})
}

func ProfHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		var u User
		if err := db.First(&u, uid).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"id":       u.ID,
				"username": u.Usr,
				"role":     u.Role,
			},
		})
	}
}

func ChgPwdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		var r ChgPwdReq
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var u User
		if err := db.First(&u, uid).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(u.Pwd), []byte(r.CurPwd)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Current password is incorrect"})
			return
		}

		if err := ValidatePwd(r.NewPwd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hp, err := bcrypt.GenerateFromPassword([]byte(r.NewPwd), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		u.Pwd = string(hp)
		if err := db.Save(&u).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie(
			"auth_token",
			"",
			-1, // (negative value to expire immediately)
			"/",
			"",
			true,
			true,
		)
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged out successfully",
		})
	}
}
