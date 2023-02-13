package initializers

import "github.com/jcblastor/restapi-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
