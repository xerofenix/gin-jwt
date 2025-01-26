package initializers

import (
	"github.com/xerofenix/gin-jwt/models"
)

func SyncDatabse() {
	DB.AutoMigrate(&models.User{})
}
