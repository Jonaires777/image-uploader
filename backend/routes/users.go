package routes

import (
	"net/http"

	"github.com/Jonaires777/image-uploader/db"
	"github.com/Jonaires777/image-uploader/models"
	"github.com/Jonaires777/image-uploader/models/dtos"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func registerUser(context *gin.Context) {
	var input dtos.UserDTO

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  input.Password,
	}

	if err := user.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.HashPassword(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":        user.ID,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"email":     user.Email,
		},
	})
}

func loginUser(context *gin.Context) {
	var input dtos.UserLoginDTO

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	row := db.DB.QueryRow("SELECT id, firstname, lastname, email, password FROM users WHERE email = $1", input.Email)
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":        user.ID,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"email":     user.Email,
		},
	})
}
