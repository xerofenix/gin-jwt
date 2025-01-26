package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xerofenix/gin-jwt/initializers"
	"github.com/xerofenix/gin-jwt/models"
)

func RequireAuth(c *gin.Context) {

	//get the cookies of req
	tockenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithError(401, err)
	}

	//decode/validate it
	tocken, err := jwt.Parse(tockenString, func(tocken *jwt.Token) (interface{}, error) {

		if _, ok := tocken.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", tocken.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := tocken.Claims.(jwt.MapClaims); ok && tocken.Valid {

		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(401, gin.H{
				"messsage": "tocken expired",
			})
		}

		//find the user with tocken sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "user not found with given tocken",
			})
		}

		//attach to request
		c.Set("user", user)

		//continue
		c.Next()
	} else {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "tocken is not valid",
		})
	}
}
