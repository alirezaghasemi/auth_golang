package handlers

import (
	"net/http"
	"time"

	"auth/db"
	"auth/models"
	"auth/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Credentials struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8"`
}

func Register(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		utils.Logger.Println("Invalid request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validate.Struct(creds); err != nil {
		utils.Logger.Println("Validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.Println("Failed to hash password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Username: creds.Username, Password: string(hashedPassword)}
	if err := db.DB.Create(&user).Error; err != nil {
		utils.Logger.Println("Failed to create user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	utils.Logger.Println("User registered successfully:", user.Username)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		utils.Logger.Println("Invalid request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validate.Struct(creds); err != nil {
		utils.Logger.Println("Validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		utils.Logger.Println("User not found:", creds.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		utils.Logger.Println("Incorrect password for user:", creds.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(utils.JwtKey)
	if err != nil {
		utils.Logger.Println("Failed to generate token for user:", user.Username)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	utils.Logger.Println("User logged in successfully:", user.Username)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
