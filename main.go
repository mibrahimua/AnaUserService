package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"strings"
)

type User struct {
	ID       int
	Email    string
	Phone    string
	Password string // In practice, hash the password before storing it
}

var users []User

func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define your routes
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// @Summary Register a new user
// @Description Register a new user with email and phone
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User object for registration"
// @Success 201 {object} string "User registered successfully"
// @Failure 400 {object} string "Invalid input data"
// @Failure 409 {object} string "Email or phone number already registered"
// @Router /register [post]
func registerHandler(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Check if the email or phone is already registered
	if isEmailTaken(newUser.Email) || isPhoneTaken(newUser.Phone) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email or phone number already registered"})
		return
	}

	// In a real application, you should hash the password before storing it
	// newUser.Password = hashPassword(newUser.Password)

	// Assign a unique ID (e.g., increment a counter)
	newUser.ID = len(users) + 1

	// Append the new user to the users list
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func loginHandler(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Find the user by email or phone
	user := findUserByEmailOrPhone(loginRequest.Email, loginRequest.Phone)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// In practice, you should verify the hashed password
	// if !verifyPassword(loginRequest.Password, user.Password) {
	//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
	//     return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func isEmailTaken(email string) bool {
	for _, user := range users {
		if strings.EqualFold(user.Email, email) {
			return true
		}
	}
	return false
}

func isPhoneTaken(phone string) bool {
	for _, user := range users {
		if user.Phone == phone {
			return true
		}
	}
	return false
}

func findUserByEmailOrPhone(email, phone string) *User {
	for _, user := range users {
		if strings.EqualFold(user.Email, email) || user.Phone == phone {
			return &user
		}
	}
	return nil
}

func hashPassword(password string) string {
	// Implement password hashing here (e.g., using bcrypt)
	return password
}

func verifyPassword(inputPassword, storedPassword string) bool {
	// Implement password verification here (e.g., using bcrypt.CompareHashAndPassword)
	return inputPassword == storedPassword
}
