package initializers

import "goJwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})

}
