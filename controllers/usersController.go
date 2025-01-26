package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xerofenix/gin-jwt/initializers"
	"github.com/xerofenix/gin-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//get the email & pass of request body
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error":   "failed to read body",
			"message": err.Error(),
		})
		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "faild to hash the password",
		})
		return
	}

	//create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "faild to create user",
		})
		return
	}

	//respond
	c.JSON(200, gin.H{
		"success": "user created",
	})
}

func Login(c *gin.Context) {
	//get the email and pss of req body
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error":   "failed to read body",
			"message": err.Error(),
		})
		return
	}

	//lookup requested user
	var user models.User
	initializers.DB.First(&user, "email =?", body.Email)
	if user.ID == 0 {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//compare sent pass with saved pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//generate new tocken
	tocken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//sign and get the complete encoded tocken as a string using the secret
	tockenString, err := tocken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(400, gin.H{
			"error":   "failed to create tocken",
			"message": err.Error(),
		})
		return
	}

	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tockenString, 3600*24*30, "", "", false, true)
	c.JSON(200, gin.H{
		"tocken": tockenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	println(user.(models.User).Email)
	c.JSON(200, gin.H{
		"mesaage": user,
	})
}
